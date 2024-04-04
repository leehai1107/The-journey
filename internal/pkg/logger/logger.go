package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

var (
	sg   *zap.SugaredLogger
	zl   LoggerInterface
	once sync.Once
)

//nolint:gochecknoinits
func init() {
	zl = &EmptyLogger{}
	logger := zap.New(zapcore.NewNopCore())
	sg = logger.Sugar()
}

// Initialize logger, once
func Initialize(production bool) {
	once.Do(func() {
		var logger *zap.Logger
		if production {
			logger = zap.New(
				zapcore.NewCore(
					zapcore.NewJSONEncoder(zapcore.EncoderConfig{
						TimeKey:        "ts",
						LevelKey:       "level",
						NameKey:        "logger",
						CallerKey:      "caller",
						MessageKey:     "msg",
						StacktraceKey:  "stacktrace",
						LineEnding:     "\r\n", // Splunk requires \r\n
						EncodeLevel:    LowercaseLevelEncoder,
						EncodeTime:     zapcore.EpochTimeEncoder,
						EncodeDuration: zapcore.SecondsDurationEncoder,
						EncodeCaller:   zapcore.ShortCallerEncoder,
					}),
					os.Stdout,
					zap.InfoLevel))
		} else {
			config := zap.NewDevelopmentConfig()
			config.EncoderConfig.EncodeLevel = CapitalLevelEncoder
			config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(timeLayout)
			logger, _ = config.Build()
		}
		// Add option: AddCallerSkip to skip report wrapper as caller in log file
		logger = logger.WithOptions(zap.AddCallerSkip(1))
		sg = logger.Sugar()
		zl = sg
	})
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	zl.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	zl.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	zl.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	zl.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then paniczl. (See DPanicLevel for detailzl.)
func DPanic(args ...interface{}) {
	zl.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then paniczl.
func Panic(args ...interface{}) {
	zl.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls ozl.Exit.
func Fatal(args ...interface{}) {
	zl.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	zl.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	zl.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	zl.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	zl.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then paniczl. (See DPanicLevel for detailzl.)
func DPanicf(template string, args ...interface{}) {
	zl.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then paniczl.
func Panicf(template string, args ...interface{}) {
	zl.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls ozl.Exit.
func Fatalf(template string, args ...interface{}) {
	zl.Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	zl.With(keysAndValues...).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	zl.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	zl.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	zl.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	zl.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then paniczl. (See DPanicLevel for detailzl.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	zl.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then paniczl. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	zl.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls ozl.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	zl.Fatalw(msg, keysAndValues...)
}

// Sync flushes any buffered log entriezl.
func Sync() error {
	return zl.Sync()
}

func SugaredLogger() *zap.SugaredLogger {
	return sg
}
