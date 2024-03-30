package logger

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// String returns a lower-case ASCII representation of the log level.
func LevelString(l zapcore.Level) string {
	switch l {
	case zapcore.DebugLevel:
		return "[debug]"
	case zapcore.InfoLevel:
		return "[info]"
	case zapcore.WarnLevel:
		return "[warn]"
	case zapcore.ErrorLevel:
		return "[error]"
	case zapcore.DPanicLevel:
		return "[dpanic]"
	case zapcore.PanicLevel:
		return "[panic]"
	case zapcore.FatalLevel:
		return "[fatal]"
	default:
		return fmt.Sprintf("[Level(%d)]", l)
	}
}

// CapitalString returns an all-caps ASCII representation of the log level.
func LevelCapitalString(l zapcore.Level) string {
	// Printing levels in all-caps is common enough that we should export this
	// functionality.
	switch l {
	case zapcore.DebugLevel:
		return "[DEBUG]"
	case zapcore.InfoLevel:
		return "[INFO]"
	case zapcore.WarnLevel:
		return "[WARN]"
	case zapcore.ErrorLevel:
		return "[ERROR]"
	case zapcore.DPanicLevel:
		return "[DPANIC]"
	case zapcore.PanicLevel:
		return "[PANIC]"
	case zapcore.FatalLevel:
		return "[FATAL]"
	default:
		return fmt.Sprintf("[LEVEL(%d)]", l)
	}
}
