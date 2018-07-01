package web

import (
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
)

// RecoveryMiddleware is a Server middleware that recovers from any panics and writes a 500 if there was one.
type RecoveryMiddleware struct {
	log              Logger
	PrintStack       bool
	ErrorHandlerFunc func(interface{})
	StackAll         bool
	StackSize        int
}

// NewRecovery returns a new instance of RecoveryMiddleware
func NewRecovery(log Logger) *RecoveryMiddleware {
	return &RecoveryMiddleware{
		log:        log,
		PrintStack: true,
		StackAll:   false,
		StackSize:  1024 * 8,
	}
}

func (self *RecoveryMiddleware) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	defer func() {
		if e := recover(); e != nil {
			if rw.Header().Get("Content-Type") == "" {
				rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
			}

			rw.WriteHeader(http.StatusInternalServerError)

			stack := make([]byte, self.StackSize)
			stack = stack[:runtime.Stack(stack, self.StackAll)]

			f := "PANIC: %s\n%s"
			self.log.Error(f, e, stack)

			if self.PrintStack {
				fmt.Fprintf(rw, f, e, stack)
			}

			if self.ErrorHandlerFunc != nil {
				func() {
					defer func() {
						if e := recover(); e != nil {
							self.log.Error("provided ErrorHandlerFunc panic'd: %s, trace:\n%s", e, debug.Stack())
							self.log.Error("%s\n", debug.Stack())
						}
					}()
					self.ErrorHandlerFunc(e)
				}()
			}
		}
	}()

	next(rw, req)
}
