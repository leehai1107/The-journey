package logger

type LoggerInterface interface {
	// Debug uses fmt.Sprint to construct and log a message.
	Debug(args ...interface{})

	// Info uses fmt.Sprint to construct and log a message.
	Info(args ...interface{})

	// Warn uses fmt.Sprint to construct and log a message.
	Warn(args ...interface{})

	// Error uses fmt.Sprint to construct and log a message.
	Error(args ...interface{})

	// DPanic uses fmt.Sprint to construct and log a message. In development, the
	// logger then paniczl. (See DPanicLevel for detailzl.)
	DPanic(args ...interface{})

	// Panic uses fmt.Sprint to construct and log a message, then paniczl.
	Panic(args ...interface{})

	// Fatal uses fmt.Sprint to construct and log a message, then calls ozl.Exit.
	Fatal(args ...interface{})

	// Debugf uses fmt.Sprintf to log a templated message.
	Debugf(template string, args ...interface{})

	// Infof uses fmt.Sprintf to log a templated message.
	Infof(template string, args ...interface{})

	// Warnf uses fmt.Sprintf to log a templated message.
	Warnf(template string, args ...interface{})

	// Errorf uses fmt.Sprintf to log a templated message.
	Errorf(template string, args ...interface{})

	// DPanicf uses fmt.Sprintf to log a templated message. In development, the
	// logger then paniczl. (See DPanicLevel for detailzl.)
	DPanicf(template string, args ...interface{})

	// Panicf uses fmt.Sprintf to log a templated message, then paniczl.
	Panicf(template string, args ...interface{})

	// Fatalf uses fmt.Sprintf to log a templated message, then calls ozl.Exit.
	Fatalf(template string, args ...interface{})

	// Debugw logs a message with some additional context. The variadic key-value
	// pairs are treated as they are in With.
	//
	// When debug-level logging is disabled, this is much faster than
	//  zl.With(keysAndValues...).Debug(msg)
	Debugw(msg string, keysAndValues ...interface{})

	// Infow logs a message with some additional context. The variadic key-value
	// pairs are treated as they are in With.
	Infow(msg string, keysAndValues ...interface{})

	// Warnw logs a message with some additional context. The variadic key-value
	// pairs are treated as they are in With.
	Warnw(msg string, keysAndValues ...interface{})

	// Errorw logs a message with some additional context. The variadic key-value
	// pairs are treated as they are in With.
	Errorw(msg string, keysAndValues ...interface{})

	// DPanicw logs a message with some additional context. In development, the
	// logger then paniczl. (See DPanicLevel for detailzl.) The variadic key-value
	// pairs are treated as they are in With.
	DPanicw(msg string, keysAndValues ...interface{})

	// Panicw logs a message with some additional context, then paniczl. The
	// variadic key-value pairs are treated as they are in With.
	Panicw(msg string, keysAndValues ...interface{})

	// Fatalw logs a message with some additional context, then calls ozl.Exit. The
	// variadic key-value pairs are treated as they are in With.
	Fatalw(msg string, keysAndValues ...interface{})

	// Sync flushes any buffered log entriezl.
	Sync() error
}
