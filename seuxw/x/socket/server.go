package socket

import (
	"context"
	"encoding/binary"
	"net/http"
	"runtime/debug"
	"seuxw/x/logger"
	"seuxw/x/socket/websocket"
	"strings"
	"sync"
	"time"
)

type Server struct {
	opts   options
	index  int // use int64 or overflow
	hs     *http.Server
	conns  map[int]*websocket.Conn
	mux    sync.RWMutex
	cond   *sync.Cond
	ctx    context.Context
	cancel context.CancelFunc
}

// NewServer creates a server which has no service registered and has not
// started to accept connection yet.
func NewServer(opt ...ServerOption) *Server {
	var opts options
	opts.minMsgSize = defaultMinMsgSize
	opts.maxMsgSize = defaultMaxMsgSize
	for _, o := range opt {
		o(&opts)
	}

	if opts.processor == nil {
		panic("server must be set a processor")
	}

	if opts.log == nil {
		opts.log = logger.NewStdLogger(true, true, true, true, true)
	}

	if opts.authenticator == nil {
		opts.authenticator = new(defaultAuthenticator)
	}

	s := &Server{
		opts:  opts,
		conns: make(map[int]*websocket.Conn),
	}

	s.cond = sync.NewCond(&s.mux)
	s.ctx, s.cancel = context.WithCancel(context.Background())

	return s
}

// Serve accepts incoming connections on the listener
func (s *Server) Serve() {
	httpmux := http.NewServeMux()
	httpmux.Handle(s.opts.url, websocket.Handler(s.handleConn))

	s.hs = &http.Server{
		Addr:           s.opts.addr,
		Handler:        httpmux,
		ReadTimeout:    time.Duration(s.opts.readTimeout) * time.Second,
		MaxHeaderBytes: 1024 * 4,
		//ErrorLog:      opts.log,
	}

	if e := s.hs.ListenAndServe(); e != nil {
		s.opts.log.Fatal("start net server failed: %v", e)
	}
}

func (s *Server) handleConn(ws *websocket.Conn) {
	if ws.Request().ParseForm() != nil {
		return
	}

	form := ws.Request().Form
	user, cer, ip := form.Get("id"), form.Get("cer"), strings.Split(ws.Request().RemoteAddr, ":")[0]
	if e := s.opts.authenticator.Auth(user, cer); e != nil {
		s.opts.log.Trace("connection %s-%s@%s auth faild", user, cer, ip)
		return
	}

	id := s.addConn(ws)
	s.opts.log.Trace("connection %d-%s@%s connected", id, user, ip)

	defer func() {
		// callback
		s.opts.processor.OnDisconnected((*Connection)(ws))

		s.removeConn(id)
		s.opts.log.Trace("connection %d-%s@%s disconnected", id, user, ip)

		if e := recover(); e != nil {
			s.opts.log.Trace("panic in connection processor: %s: %s", e, debug.Stack())
		}
	}()

	// callback
	s.opts.processor.OnConnected((*Connection)(ws))

	for {
		var buf []byte
		//socket.SetReadDeadline(time.Now().Add(time.Minute*5))
		if e := websocket.Message.Receive(ws, &buf); e != nil {
			break
		}

		if length := len(buf); length < s.opts.minMsgSize || length > s.opts.maxMsgSize {
			s.opts.log.Trace("connection[%d-%s] invalid message length: %d", id, user, length)
			break
		}

		msg := &Message{
			SN:      int(binary.BigEndian.Uint32(buf)),
			ID:      int(binary.BigEndian.Uint32(buf[4:])),
			RawData: buf[8:],
			conn:    ws,
		}
		// callback
		s.opts.processor.OnMessage(msg)
	}
}

func (s *Server) addConn(conn *websocket.Conn) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.conns != nil {
		s.index++
		s.conns[s.index] = conn
		return s.index
	}
	return 0
}

func (s *Server) removeConn(id int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.conns != nil {
		delete(s.conns, id)
		s.cond.Broadcast()
	}
}

// Stop stops the server. It immediately closes all open
// connections and listeners.
func (s *Server) Stop() {
	s.mux.Lock()
	st := s.conns
	s.conns = nil
	// interrupt GracefulStop if Stop and GracefulStop are called concurrently.
	s.cond.Broadcast()
	s.mux.Unlock()

	s.hs.Close()

	for _, c := range st {
		c.Close()
	}

	s.mux.Lock()
	s.cancel()
	// if s.events != nil {
	// 	s.events.Finish()
	// 	s.events = nil
	// }
	s.mux.Unlock()
}
