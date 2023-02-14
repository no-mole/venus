package logger

import (
	"github.com/hashicorp/go-hclog"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
)

func New(name string, level hclog.Level, zapLogger *zap.Logger) hclog.Logger {
	return &logger{
		name:      name,
		level:     level,
		zapLogger: zapLogger.Named(name),
	}
}

type logger struct {
	name      string
	level     hclog.Level
	zapLogger *zap.Logger
	fields    []zap.Field
}

func (l *logger) Log(level hclog.Level, msg string, args ...interface{}) {
	switch level {
	case hclog.Trace:
		l.Trace(msg, args...)
	case hclog.Debug:
		l.Debug(msg, args...)
	case hclog.Info:
		l.Info(msg, args...)
	case hclog.Warn:
		l.Warn(msg, args...)
	case hclog.Error:
		l.Error(msg, args...)
	default:
		l.Error(msg, args...)
	}
}

func (l *logger) Trace(msg string, args ...interface{}) {
	l.zapLogger.Debug(msg, argsToFields(args...)...)
}

func (l *logger) Debug(msg string, args ...interface{}) {
	l.zapLogger.Debug(msg, argsToFields(args...)...)

}

func (l *logger) Info(msg string, args ...interface{}) {
	l.zapLogger.Info(msg, argsToFields(args...)...)
}

func (l *logger) Warn(msg string, args ...interface{}) {
	l.zapLogger.Warn(msg, argsToFields(args...)...)
}

func (l *logger) Error(msg string, args ...interface{}) {
	l.zapLogger.Error(msg, argsToFields(args...)...)
}

func argsToFields(args ...interface{}) (fields []zap.Field) {
	var extra interface{}
	//单数
	if len(args)%2 == 1 {
		extra = args[len(args)-1]
		args = args[:len(args)-1]
	}
	for i := 0; i < len(args); i += 2 {
		str, ok := args[i].(string)
		if !ok {
			continue
		}
		fields = append(fields, zap.Any(str, args[i+1]))
	}
	if extra != nil {
		fields = append(fields, zap.Any("extra", extra))
	}
	return
}

func (l *logger) IsTrace() bool {
	return l.level == hclog.Trace
}

func (l *logger) IsDebug() bool {
	return l.level == hclog.Debug
}

func (l *logger) IsInfo() bool {
	return l.level == hclog.Info
}

func (l *logger) IsWarn() bool {
	return l.level == hclog.Warn
}

func (l *logger) IsError() bool {
	return l.level == hclog.Error
}

func (l *logger) ImpliedArgs() []interface{} {
	return []interface{}{}
}

func (l *logger) With(args ...interface{}) hclog.Logger {
	return &logger{
		name:      l.name,
		level:     l.level,
		zapLogger: l.zapLogger.Named(l.name),
		fields:    append(argsToFields(args...), l.fields...),
	}
}

func (l *logger) Name() string {
	return l.name
}

func (l *logger) Named(name string) hclog.Logger {
	return New(name, l.level, l.zapLogger)
}

func (l *logger) ResetNamed(name string) hclog.Logger {
	return l.Named(name)
}

func (l *logger) SetLevel(level hclog.Level) {
	l.level = level
}

func (l *logger) StandardLogger(opts *hclog.StandardLoggerOptions) *log.Logger {
	return log.New(l.StandardWriter(opts), "", log.LstdFlags)
}

func (l *logger) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	return os.Stderr
}

func (l *logger) mergeArgs(args ...interface{}) []zap.Field {
	return append(argsToFields(args), l.fields...)
}
