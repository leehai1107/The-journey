package logger

type EmptyLogger struct {
}

// Debug uses fmt.Sprint to construct and log a message.
func (zl *EmptyLogger) Debug(args ...interface{}) {
}

// Info uses fmt.Sprint to construct and log a message.
func (zl *EmptyLogger) Info(args ...interface{}) {
}

// Warn uses fmt.Sprint to construct and log a message.
func (zl *EmptyLogger) Warn(args ...interface{}) {
}

// Error uses fmt.Sprint to construct and log a message.
func (zl *EmptyLogger) Error(args ...interface{}) {
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then paniczl. (See DPanicLevel for detailzl.)
func (zl *EmptyLogger) DPanic(args ...interface{}) {
}

// Panic uses fmt.Sprint to construct and log a message, then paniczl.
func (zl *EmptyLogger) Panic(args ...interface{}) {
}

// Fatal uses fmt.Sprint to construct and log a message, then calls ozl.Exit.
func (zl *EmptyLogger) Fatal(args ...interface{}) {
}

// Debugf uses fmt.Sprintf to log a templated message.
func (zl *EmptyLogger) Debugf(template string, args ...interface{}) {
}

// Infof uses fmt.Sprintf to log a templated message.
func (zl *EmptyLogger) Infof(template string, args ...interface{}) {
}

// Warnf uses fmt.Sprintf to log a templated message.
func (zl *EmptyLogger) Warnf(template string, args ...interface{}) {
}

// Errorf uses fmt.Sprintf to log a templated message.
func (zl *EmptyLogger) Errorf(template string, args ...interface{}) {
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then paniczl. (See DPanicLevel for detailzl.)
func (zl *EmptyLogger) DPanicf(template string, args ...interface{}) {
}

// Panicf uses fmt.Sprintf to log a templated message, then paniczl.
func (zl *EmptyLogger) Panicf(template string, args ...interface{}) {
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls ozl.Exit.
func (zl *EmptyLogger) Fatalf(template string, args ...interface{}) {
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	zl.With(keysAndValues...).Debug(msg)
func (zl *EmptyLogger) Debugw(msg string, keysAndValues ...interface{}) {
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (zl *EmptyLogger) Infow(msg string, keysAndValues ...interface{}) {
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (zl *EmptyLogger) Warnw(msg string, keysAndValues ...interface{}) {
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (zl *EmptyLogger) Errorw(msg string, keysAndValues ...interface{}) {
}

// DPanicw logs a message with some additional context. In development, the
// logger then paniczl. (See DPanicLevel for detailzl.) The variadic key-value
// pairs are treated as they are in With.
func (zl *EmptyLogger) DPanicw(msg string, keysAndValues ...interface{}) {
}

// Panicw logs a message with some additional context, then paniczl. The
// variadic key-value pairs are treated as they are in With.
func (zl *EmptyLogger) Panicw(msg string, keysAndValues ...interface{}) {
}

// Fatalw logs a message with some additional context, then calls ozl.Exit. The
// variadic key-value pairs are treated as they are in With.
func (zl *EmptyLogger) Fatalw(msg string, keysAndValues ...interface{}) {
}

func (zl *EmptyLogger) Printf(format string, items ...interface{}) {
}

// Sync flushes any buffered log entriezl.
func (zl *EmptyLogger) Sync() error {
	return nil
}
