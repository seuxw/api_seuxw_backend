package socket

var (
	defaultMinMsgSize = 16       // use 16Byte as the default message mimimum size
	defaultMaxMsgSize = 1024 * 2 // use 2K as the default message maximum size
)

type options struct {
	addr              string
	url               string
	minMsgSize        int
	maxMsgSize        int
	readTimeout       int
	maxConcurrentConn int
	log               Logger
	processor         Processor
	authenticator     Authenticator
}

// A ServerOption sets options.
type ServerOption func(*options)

// CustomLogger returns a ServerOption that sets a log for server.
func CustomLogger(l Logger) ServerOption {
	return func(o *options) {
		o.log = l
	}
}

// CustomAuthenticator returns a ServerOption that sets authenticator for server.
func CustomAuthenticator(a Authenticator) ServerOption {
	return func(o *options) {
		o.authenticator = a
	}
}

// Process returns a ServerOption that sets a processor for connection handler.
func Process(p Processor) ServerOption {
	return func(o *options) {
		o.processor = p
	}
}

func Addr(s string) ServerOption {
	return func(o *options) {
		o.addr = s
	}
}

func URL(s string) ServerOption {
	return func(o *options) {
		o.url = s
	}
}

// MinMsgSize returns a ServerOption to set the max message size in bytes for inbound mesages.
// If this is not set, abc uses the default 18Byte.
func MinMsgSize(m int) ServerOption {
	return func(o *options) {
		o.minMsgSize = m
	}
}

// MaxMsgSize returns a ServerOption to set the max message size in bytes for inbound mesages.
// If this is not set, abc uses the default 2K.
func MaxMsgSize(m int) ServerOption {
	return func(o *options) {
		o.maxMsgSize = m
	}
}

func ReadTimeout(t int) ServerOption {
	return func(o *options) {
		o.readTimeout = t
	}
}

// MaxConcurrentConn returns a ServerOption that will apply a limit on the number
// of concurrent connection to the Server.
func MaxConcurrentConn(n int) ServerOption {
	return func(o *options) {
		o.maxConcurrentConn = n
	}
}
