package web

import (
	"net/http"
	"seuxw/x/web/mux"
)

// Handler handler is an interface that objects can implement to be registered to serve as middleware
// in the Server middleware stack.
// ServeHTTP should yield to the next middleware in the chain by invoking the next http.HandlerFunc
// passed in.
//
// If the Handler writes to the ResponseWriter, the next http.HandlerFunc should not be invoked.
type Handler interface {
	ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

// HandlerFunc is an adapter to allow the use of ordinary functions as Server handlers.
// If f is a function with the appropriate signature, HandlerFunc(f) is a Handler object that calls f.
type HandlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (h HandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h(rw, r, next)
}

type Middleware struct {
	handler Handler
	next    *Middleware
}

func (m Middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	m.handler.ServeHTTP(rw, r, m.next.ServeHTTP)
}

// Wrap converts a http.Handler into a web.Handler so it can be used as a Server
// middleware. The next http.HandlerFunc is automatically called after the Handler
// is executed.
func Wrap(handler http.Handler) Handler {
	return HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		handler.ServeHTTP(rw, r)
		next(rw, r)
	})
}

// Server is a stack of Middleware Handlers that can be invoked as an http.Handler.
// Server middleware is evaluated in the order that they are added to the stack using
// the Filter and UseHandler methods.
type Server struct {
	*mux.Router
	middleware Middleware
	handlers   []Handler
	log        Logger
}

// New returns a new Server instance with no middleware preconfigured.
func New(log Logger, handlers ...Handler) *Server {
	return &Server{
		Router:     mux.NewRouter(),
		handlers:   handlers,
		middleware: build(handlers),
		log:        log,
	}
}

// NewServer returns a new Server instance with the default middleware already
// in the stack.
//
// Recovery - Panic Recovery Middleware
// Logger - Request/Response Logging
// Static - Static File Serving
func NewServer(log Logger) *Server {
	return New(log, NewRecovery(log), NewLogger(log))
}

// Run is a convenience function that runs the server stack as an HTTP
// server. The addr string takes the same format as http.ListenAndServe.
func (self *Server) Serve(addr string) {
	self.log.Trace("webserver serve on: %s", addr)
	self.UseHandler(self.Router)
	self.log.Fatal("%v", http.ListenAndServe(addr, self))
}

func (self *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	self.middleware.ServeHTTP(NewResponseWriter(rw), r)
}

// Filter adds a Handler onto the middleware stack. Handlers are invoked in the order they are added to a Server.
func (self *Server) Filter(handler Handler) *Server {
	if handler == nil {
		panic("handler cannot be nil")
	}

	self.handlers = append(self.handlers, handler)
	self.middleware = build(self.handlers)
	return self
}

// FilterFunc adds a Server-style handler function onto the middleware stack.
func (self *Server) FilterFunc(handlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) *Server {
	self.Filter(HandlerFunc(handlerFunc))
	return self
}

// UseHandler adds a http.Handler onto the middleware stack. Handlers are invoked in the order they are added to a Server.
func (self *Server) UseHandler(handler http.Handler) {
	self.Filter(Wrap(handler))
}

// UseHandler adds a http.HandlerFunc-style handler function onto the middleware stack.
func (self *Server) UseHandlerFunc(handlerFunc func(rw http.ResponseWriter, r *http.Request)) {
	self.UseHandler(http.HandlerFunc(handlerFunc))
}

func build(handlers []Handler) Middleware {
	var next Middleware

	if len(handlers) == 0 {
		return voidMiddleware()
	} else if len(handlers) > 1 {
		next = build(handlers[1:])
	} else {
		next = voidMiddleware()
	}

	return Middleware{handlers[0], &next}
}

func voidMiddleware() Middleware {
	return Middleware{
		HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}),
		&Middleware{},
	}
}
