package log

import (
	"fmt"
	"runtime"
	"strings"

	"ig.engagements/libs/host"
	"github.com/sirupsen/logrus"
)

type Entry interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Panic(format string, args ...interface{})

	WithError(err error) Entry
	WithField(key string, value interface{}) Entry
	WithFields(fields map[string]interface{}) Entry
}

type entry struct {
	l  *logger
	e  *logrus.Entry
	sl int
}

func (e *entry) log(level Level, format string, args ...interface{}) {
	var msg string
	if len(args) == 0 {
		msg = format
	}
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	entry := e.e
	stackLevel := e.sl
	if stackLevel == -1 {
		stackLevel = e.l.stackLevel
	}
	// Append caller line number if possible
	_, file, line, ok := runtime.Caller(stackLevel)
	parts := strings.Split(file, splitter)
	if ok && len(parts) > 1 {
		slashIndex := strings.Index(parts[1], "/")
		entry = entry.WithField("location", fmt.Sprintf("%s:%d", parts[1][slashIndex+1:], line))
	}

	// Append host identifier to the log.
	h := host.Host()
	if len(h) > 0 {
		entry = entry.WithField("host", host.Host())
	}

	switch level {
	case LevelDebug:
		entry.Debug(msg)
	case LevelInfo:
		entry.Info(msg)
	case LevelWarn:
		entry.Warn(msg)
	case LevelError:
		entry.Error(msg)
	case LevelPanic:
		entry.Panic(msg)
	default:
		entry.Debug(msg)
	}
}

func (e *entry) Debug(format string, args ...interface{}) {
	e.log(LevelDebug, format, args...)
}

func (e *entry) Info(format string, args ...interface{}) {
	e.log(LevelInfo, format, args...)
}

func (e *entry) Warning(format string, args ...interface{}) {
	e.log(LevelWarn, format, args...)
}

func (e *entry) Error(format string, args ...interface{}) {
	e.log(LevelError, format, args...)
}

func (e *entry) Panic(format string, args ...interface{}) {
	e.log(LevelPanic, format, args...)
}

func (e *entry) WithError(err error) Entry {
	return &entry{
		e.l,
		e.e.WithError(err),
		-1,
	}
}

func (e *entry) WithField(key string, value interface{}) Entry {
	return &entry{
		e.l,
		e.e.WithField(key, value),
		-1,
	}
}

func (e *entry) WithFields(fields map[string]interface{}) Entry {
	return &entry{
		e.l,
		e.e.WithFields(fields),
		-1,
	}
}
