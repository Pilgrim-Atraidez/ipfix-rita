package logging

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//Fields holds arbitrary data for structured logging
type Fields map[string]interface{}

//Logger provides the means to log errors, status updates, etc.
type Logger interface {
	Error(err error, fields Fields)
	Warn(msg string, fields Fields)
	Info(msg string, fields Fields)
}

//stackTracer is defined as per https://godoc.org/github.com/pkg/errors#Cause
type stackTracer interface {
	StackTrace() errors.StackTrace
}

//causer is defined as per https://godoc.org/github.com/pkg/errors#Cause
type causer interface {
	Cause() error
}

//logrusLogger implements Logger by value (since its just a pointer)
type logrusLogger struct {
	logger *log.Logger
}

//NewLogrusLogger creates a Logger backed by the logrus logging system
func NewLogrusLogger() Logger {
	log := logrusLogger{
		logger: log.New(),
	}

	return log
}

func (l logrusLogger) Error(err error, fields Fields) {
	if err, ok := err.(stackTracer); ok {
		if fields == nil {
			fields = make(Fields)
		}
		fields["stacktrace"] = err.StackTrace()
	}
	if fields != nil {
		l.logger.WithFields(log.Fields(fields)).Error(err.Error())
	} else {
		l.logger.Error(err.Error())
	}
}

func (l logrusLogger) Info(msg string, fields Fields) {
	if fields != nil {
		l.logger.WithFields(log.Fields(fields)).Info(msg)
	} else {
		l.logger.Info(msg)
	}
}

func (l logrusLogger) Warn(msg string, fields Fields) {
	if fields != nil {
		l.logger.WithFields(log.Fields(fields)).Warn(msg)
	} else {
		l.logger.Warn(msg)
	}
}
