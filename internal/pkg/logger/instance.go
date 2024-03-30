package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"go.uber.org/zap"
)

const (
	callerSkip = 4
)

type LogInstance struct {
	zl             LoggerInterface
	withCallerInfo bool
}

func NewLogInstance(zl LoggerInterface, withCallerInfo bool) LoggerInterface {
	return &LogInstance{zl, withCallerInfo}
}

func (l *LogInstance) formatCaller() string {
	fpcs := make([]uintptr, 1)
	n := runtime.Callers(callerSkip, fpcs)
	if n == 0 {
		return "<nocaller>"
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		return "<nocaller>"
	}

	file, line := caller.FileLine(fpcs[0] - 1)
	_, fileName := filepath.Split(file)
	return fmt.Sprintf("%s@%s:%d", l.getFunctionName(caller.Name()), fileName, line)
}

func (l *LogInstance) getFunctionName(fullFunctionPath string) string {
	items := strings.Split(fullFunctionPath, ".")
	if len(items) > 0 {
		return items[len(items)-1]
	}
	return fullFunctionPath
}

func (l *LogInstance) withCaller() LoggerInterface {
	sgr, ok := l.zl.(*zap.SugaredLogger)
	if l.withCallerInfo && ok {
		return sgr.With("caller", l.formatCaller())
	} else {
		return l.zl
	}
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *LogInstance) Debug(args ...interface{}) {
	l.withCaller().Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *LogInstance) Info(args ...interface{}) {
	l.withCaller().Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *LogInstance) Warn(args ...interface{}) {
	l.withCaller().Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *LogInstance) Error(args ...interface{}) {
	l.withCaller().Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then paniczl. (See DPanicLevel for detailzl.)
func (l *LogInstance) DPanic(args ...interface{}) {
	l.withCaller().DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then paniczl.
func (l *LogInstance) Panic(args ...interface{}) {
	l.withCaller().Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls ozl.Exit.
func (l *LogInstance) Fatal(args ...interface{}) {
	l.withCaller().Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *LogInstance) Debugf(template string, args ...interface{}) {
	l.withCaller().Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *LogInstance) Infof(template string, args ...interface{}) {
	l.withCaller().Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *LogInstance) Warnf(template string, args ...interface{}) {
	l.withCaller().Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *LogInstance) Errorf(template string, args ...interface{}) {
	l.withCaller().Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then paniczl. (See DPanicLevel for detailzl.)
func (l *LogInstance) DPanicf(template string, args ...interface{}) {
	l.withCaller().DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then paniczl.
func (l *LogInstance) Panicf(template string, args ...interface{}) {
	l.withCaller().Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls ozl.Exit.
func (l *LogInstance) Fatalf(template string, args ...interface{}) {
	l.withCaller().Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	zl.With(keysAndValues...).Debug(msg)
func (l *LogInstance) Debugw(msg string, keysAndValues ...interface{}) {
	l.withCaller().Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *LogInstance) Infow(msg string, keysAndValues ...interface{}) {
	l.withCaller().Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *LogInstance) Warnw(msg string, keysAndValues ...interface{}) {
	l.withCaller().Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *LogInstance) Errorw(msg string, keysAndValues ...interface{}) {
	l.withCaller().Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then paniczl. (See DPanicLevel for detailzl.) The variadic key-value
// pairs are treated as they are in With.
func (l *LogInstance) DPanicw(msg string, keysAndValues ...interface{}) {
	l.withCaller().DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then paniczl. The
// variadic key-value pairs are treated as they are in With.
func (l *LogInstance) Panicw(msg string, keysAndValues ...interface{}) {
	l.withCaller().Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls ozl.Exit. The
// variadic key-value pairs are treated as they are in With.
func (l *LogInstance) Fatalw(msg string, keysAndValues ...interface{}) {
	l.withCaller().Fatalw(msg, keysAndValues...)
}

// Sync flushes any buffered log entriezl.
func (l *LogInstance) Sync() error {
	return l.zl.Sync()
}
