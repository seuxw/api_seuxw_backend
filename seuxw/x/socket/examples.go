package main

import (
	"flag"
	"seuxw/x/logger"
	"seuxw/x/socket"
)

var log = logger.NewStdLogger(true, true, true, true, true)

type server struct {
	*socket.Server
}

func (s *server) OnConnected(conn *socket.Connection) {
	log.Debug("connected: %v", conn)
}

func (s *server) OnMessage(msg *socket.Message) {
	log.Debug("message: %s, %v", msg.RawData, msg)
	//s.push(msg)
}

func (s *server) OnDisconnected(conn *socket.Connection) {
	log.Debug("disconnected: %v", conn)
}

func (s *server) serve() {
	opts := []socket.ServerOption{socket.Addr(":10001"), socket.URL("/auth"), socket.Process(s)}
	s.Server = socket.NewServer(opts...)
	s.Serve()
}

func main() {
	flag.Parse()
	log.Trace("examples: v0.0.1")

	var s server
	s.serve()
}
