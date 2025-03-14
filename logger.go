package logger

import (
	"fmt"
	"time"
)

type MinimumLogger interface {
  Infof(format string, args ...interface{})
  Errorf(format string, args ...interface{})
}

type Logger interface {
  MinimumLogger
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

var _ = MinimumLogger(&defaultLogger{})
var _ = Logger(&defaultLogger{})

var DefaultLogger = defaultLogger{
	level: LogLevelInfo,
}

func DebugLabel() string {
	return "[DEBUG]"
}

func InfoLabel() string {
	return WithColor(ColorCyan, "[INFO]")
}

func WarnLabel() string {
	return WithColor(ColorYellow, "[WARN]")
}

func ErrorLabel() string {
	return WithColor(ColorRed, "[ERROR]")
}

func DateLabel() string {
	return DateFormat(time.Now())
}

func DateFormat(t time.Time) string {
	return WithColor(ColorGray, t.Format("2006-01-02T15:04:05 -0700"))
}

type defaultLogger struct {
	level LogLevel
}

func (l *defaultLogger) SetLevel(level LogLevel) {
	l.level = level
}

func (l defaultLogger) SDebugf(format string, args ...any) string {
	return fmt.Sprintf(fmt.Sprintf("%s %s %s\n", DebugLabel(), DateLabel(), format), args...)
}

func (l defaultLogger) SInfof(format string, args ...any) string {
	return fmt.Sprintf(fmt.Sprintf("%s %s %s\n", InfoLabel(), DateLabel(), format), args...)
}

func (l defaultLogger) SErrorf(format string, args ...any) string {
	return fmt.Sprintf(fmt.Sprintf("%s %s %s\n", ErrorLabel(), DateLabel(), format), args...)
}

func (l defaultLogger) SWarnf(format string, args ...any) string {
	return fmt.Sprintf(fmt.Sprintf("%s %s %s\n", WarnLabel(), DateLabel(), format), args...)
}

func (l defaultLogger) Debugf(format string, args ...interface{}) {
	if l.level >= LogLevelInfo {
		return
	}
	fmt.Printf(l.SDebugf(format, args...))
}

func (l defaultLogger) Infof(format string, args ...interface{}) {
	if l.level > LogLevelInfo {
		return
	}
	fmt.Printf(l.SInfof(format, args...))
}

func (l defaultLogger) Warnf(format string, args ...interface{}) {
	if l.level > LogLevelWarn {
		return
	}
	fmt.Printf(l.SWarnf(format, args...))
}

func (l defaultLogger) Errorf(format string, args ...interface{}) {
	if l.level > LogLevelError {
		return
	}
	fmt.Printf(l.SErrorf(format, args...))
}

func WithColor(c, msg string) string {
	return fmt.Sprintf("%s%s%s", c, msg, ColorReset)
}

var ColorReset = "\033[0m"
var ColorRed = "\033[31m"
var ColorGreen = "\033[32m"
var ColorYellow = "\033[33m"
var ColorBlue = "\033[34m"
var ColorMagenta = "\033[35m"
var ColorCyan = "\033[36m"
var ColorGray = "\033[37m"
var ColorWhite = "\033[97m"
