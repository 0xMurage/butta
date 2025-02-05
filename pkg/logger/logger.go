package logger

import (
	"log/slog"
	"os"
)

var l = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Info(msg string, args ...any) {
	l.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	l.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	l.Error(msg, args...)
}

func Debug(msg string, args ...any) {
	l.Debug(msg, args...)
}

// Fatal is equivalent to [Info] followed by a call to [os.Exit](1).
func Fatal(msg string, args ...any) {
	l.Error(msg, args...)
	os.Exit(1)
}
