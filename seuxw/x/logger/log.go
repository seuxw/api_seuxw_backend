// Copyright 2012-2015 Apcera Inc. All rights reserved.

//Package logger provides logging facilities for the NATS server
package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger is the server logger
type Logger struct {
	logger     *log.Logger
	debug      bool
	trace      bool
	warnLabel  string
	errorLabel string
	fatalLabel string
	debugLabel string
	traceLabel string
}

// NewStdLogger creates a logger with output directed to Stderr
func NewStdLogger(time, debug, trace, colors, pid bool) *Logger {
	flags := 0
	if time {
		flags = log.LstdFlags | log.Lmicroseconds
	}

	if debug {
		flags = flags | log.Lshortfile
	}

	pre := ""
	if pid {
		pre = pidPrefix()
	}

	l := &Logger{
		logger: log.New(os.Stderr, pre, flags),
		debug:  debug,
		trace:  trace,
	}

	if colors {
		setColoredLabelFormats(l)
	} else {
		setPlainLabelFormats(l)
	}

	return l
}

// NewFileLogger creates a logger with output directed to a file
func NewFileLogger(filename string, time, debug, trace, pid bool) *Logger {
	fileflags := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	f, err := os.OpenFile(filename, fileflags, 0660)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	flags := 0
	if time {
		flags = log.LstdFlags | log.Lmicroseconds | log.Lshortfile
	}

	pre := ""
	if pid {
		pre = pidPrefix()
	}

	l := &Logger{
		logger: log.New(f, pre, flags),
		debug:  debug,
		trace:  trace,
	}

	setPlainLabelFormats(l)
	return l
}

// Generate the pid prefix string
func pidPrefix() string {
	return fmt.Sprintf("[%d] ", os.Getpid())
}

func setPlainLabelFormats(l *Logger) {
	l.debugLabel = "[DBG] "
	l.traceLabel = "[TRC] "
	l.warnLabel = "[WAR] "
	l.errorLabel = "[ERR] "
	l.fatalLabel = "[FTL] "
}

func setColoredLabelFormats(l *Logger) {
	colorFormat := "[\x1b[%dm%s\x1b[0m] "
	l.debugLabel = fmt.Sprintf(colorFormat, 36, "DBG")
	l.traceLabel = fmt.Sprintf(colorFormat, 33, "TRC")
	l.warnLabel = fmt.Sprintf(colorFormat, 32, "WAR")
	l.errorLabel = fmt.Sprintf(colorFormat, 31, "ERR")
	l.fatalLabel = fmt.Sprintf(colorFormat, 31, "FTL")
}

// Debug logs a debug statement
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.debug {
		l.logger.Printf(l.debugLabel+format, v...)
	}
}

// Trace logs a trace statement
func (l *Logger) Trace(format string, v ...interface{}) {
	if l.trace {
		l.logger.Printf(l.traceLabel+format, v...)
	}
}

// Warning logs a notice statement
func (l *Logger) Warning(format string, v ...interface{}) {
	l.logger.Printf(l.warnLabel+format, v...)
}

// Error logs an error statement
func (l *Logger) Error(format string, v ...interface{}) {
	l.logger.Printf(l.errorLabel+format, v...)
}

// Fatal logs a fatal error
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.logger.Fatalf(l.fatalLabel+format, v...)
}
