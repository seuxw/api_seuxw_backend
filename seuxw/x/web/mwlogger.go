package web

import (
	"bytes"
	"net/http"
	"seuxw/x/uuid"
	"text/template"
	"time"
)

// Logger interface
type Logger interface {
	Debug(format string, v ...interface{})
	Trace(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
}

// LogEntry is the structure,  passed to the template.
type LogEntry struct {
	Trace    string
	Status   int
	Duration time.Duration
	Method   string
	Host     string
	URL      string
	Addr     string
}

// LogEntryFormat is the format logged used by the default log instance.
var LogEntryFormat = "{{.Trace}} | {{.Status}} | {{.Duration}} | {{.Addr}} | {{.Method}} {{.Host}} {{.URL}}"

// LoggerDateFormat is the format used for date by the default log instance.
var LogDateFormat = time.RFC3339

// LogMiddleware is a middleware handler that logs the request as it goes in and the response as it goes out.
type LogMiddleware struct {
	log        Logger
	dateFormat string
	template   *template.Template
}

// NewLogger returns a new Logger instance
func NewLogger(log Logger) *LogMiddleware {
	mw := &LogMiddleware{
		log:        log,
		dateFormat: LogDateFormat,
	}
	mw.SetFormat(LogEntryFormat)
	return mw
}

func (self *LogMiddleware) SetFormat(format string) {
	self.template = template.Must(template.New("seuxw.web.log.parser").Parse(format))
}

func (self *LogMiddleware) SetDateFormat(format string) {
	self.dateFormat = format
}

func (self *LogMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	// add the uuid to request header for trace debug
	trace := uuid.New()
	r.Header.Set("X-Trace", trace)
	rw.Header().Set("Content-Type", "application/json")

	next(rw, r)

	resp := rw.(ResponseWriter)
	entry := LogEntry{
		Trace: trace,
		//StartTime: start.Format(self.dateFormat),
		Status:   resp.Status(),
		Duration: time.Since(start),
		Method:   r.Method,
		Host:     r.Host,
		URL:      r.RequestURI,
		Addr:     r.RemoteAddr,
	}

	buff := &bytes.Buffer{}
	self.template.Execute(buff, entry)
	self.log.Trace(buff.String())
}
