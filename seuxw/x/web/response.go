package web

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

// ResponseWriter is a wrapper around http.ResponseWriter that provides extra information about
// the response. It is recommended that Middleware handlers use this construct to wrap a responsewriter
// if the functionality calls for it.
type ResponseWriter interface {
	http.ResponseWriter
	http.Flusher
	// Status returns the status code of the response or 200 if the response has
	// not been written (as this is the default response code in net/http)
	Status() int
	// Written returns whether or not the ResponseWriter has been written.
	Written() bool
	// Size returns the size of the response body.
	Size() int
	// Before allows for a function to be called before the ResponseWriter has been written to. This is
	// useful for setting headers or any other operations that must happen before a response has been written.
	Before(func(ResponseWriter))
}

type beforeFunc func(ResponseWriter)

// NewResponseWriter creates a ResponseWriter that wraps an http.ResponseWriter
func NewResponseWriter(rw http.ResponseWriter) ResponseWriter {
	nrw := &responseWriter{
		ResponseWriter: rw,
	}

	if _, ok := rw.(http.CloseNotifier); ok {
		return &responseWriterCloseNotifer{nrw}
	}

	return nrw
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	size        int
	beforeFuncs []beforeFunc
}

func (self *responseWriter) WriteHeader(s int) {
	self.status = s
	self.callBefore()
	self.ResponseWriter.WriteHeader(s)
}

func (self *responseWriter) Write(b []byte) (int, error) {
	if !self.Written() {
		// The status will be StatusOK if WriteHeader has not been called yet
		self.WriteHeader(http.StatusOK)
	}
	size, err := self.ResponseWriter.Write(b)
	self.size += size
	return size, err
}

func (self *responseWriter) Status() int {
	return self.status
}

func (self *responseWriter) Size() int {
	return self.size
}

func (self *responseWriter) Written() bool {
	return self.status != 0
}

func (self *responseWriter) Before(before func(ResponseWriter)) {
	self.beforeFuncs = append(self.beforeFuncs, before)
}

func (self *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := self.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}

func (self *responseWriter) callBefore() {
	for i := len(self.beforeFuncs) - 1; i >= 0; i-- {
		self.beforeFuncs[i](self)
	}
}

func (self *responseWriter) Flush() {
	flusher, ok := self.ResponseWriter.(http.Flusher)
	if ok {
		if !self.Written() {
			// The status will be StatusOK if WriteHeader has not been called yet
			self.WriteHeader(http.StatusOK)
		}
		flusher.Flush()
	}
}

type responseWriterCloseNotifer struct {
	*responseWriter
}

func (self *responseWriterCloseNotifer) CloseNotify() <-chan bool {
	return self.ResponseWriter.(http.CloseNotifier).CloseNotify()
}
