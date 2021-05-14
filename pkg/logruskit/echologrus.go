package logruskit

import (
	"io"

	"github.com/labstack/echo/v4"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type (
	// EchoLogrus is an implementation of echo logger
	EchoLogrus struct {
		*logrus.Logger
		prefix string
	}
)

var _ echo.Logger = (*EchoLogrus)(nil)

// EchoLogger logrus logger in echo log interface
func EchoLogger(logger *logrus.Logger) *EchoLogrus {
	return &EchoLogrus{Logger: logger}
}

// Output logger output func
func (e *EchoLogrus) Output() io.Writer {
	return e.Out
}

// Prefix returns the prefix value
func (e *EchoLogrus) Prefix() string {
	return e.prefix
}

// SetPrefix set prefix of EchoLogrus
func (e *EchoLogrus) SetPrefix(prefix string) {
	e.prefix = prefix
}

// Level returns logger level
func (e *EchoLogrus) Level() log.Lvl {
	switch e.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	}
	return log.WARN
}

// SetLevel set level to logger from given log.Lvl
func (e *EchoLogrus) SetLevel(lvl log.Lvl) {
	switch lvl {
	case log.DEBUG:
		e.Logger.SetLevel(logrus.DebugLevel)
	case log.WARN:
		e.Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		e.Logger.SetLevel(logrus.ErrorLevel)
	case log.INFO:
		e.Logger.SetLevel(logrus.InfoLevel)
	default:
		logrus.Warnf("Unknown level: %v", lvl)
		e.Logger.SetLevel(logrus.WarnLevel)
	}
}

// SetHeader to set header
func (e EchoLogrus) SetHeader(h string) {}

// Printj print json log
func (e *EchoLogrus) Printj(j log.JSON) {
	e.Logger.WithFields(logrus.Fields(j)).Print()
}

// Debugj debug json log
func (e *EchoLogrus) Debugj(j log.JSON) {
	e.Logger.WithFields(logrus.Fields(j)).Debug()
}

// Infoj info json log
func (e *EchoLogrus) Infoj(j log.JSON) {
	e.Logger.WithFields(logrus.Fields(j)).Info()
}

// Warnj warning json log
func (e *EchoLogrus) Warnj(j log.JSON) {
	e.Logger.WithFields(logrus.Fields(j)).Warn()
}

// Errorj error json log
func (e *EchoLogrus) Errorj(j log.JSON) {
	e.Logger.WithFields(logrus.Fields(j)).Error()
}

// Fatalj fatal json log
func (e *EchoLogrus) Fatalj(j log.JSON) {
	e.Logger.WithFields(logrus.Fields(j)).Fatal()
}

// Panicj panic json log
func (e *EchoLogrus) Panicj(j log.JSON) {
	e.Logger.WithFields(logrus.Fields(j)).Panic()
}
