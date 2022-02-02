package logger

import (
	"context"
	"log"
	"os"
	"sync"
)

var (
	logger = log.New(os.Stderr, "", log.LstdFlags)
	logMu  sync.Mutex
)

// SetLogger set logger in outside of library
func SetLogger(l *log.Logger) {
	if l == nil {
		l = log.New(os.Stderr, "", log.LstdFlags)
	}
	logMu.Lock()
	logger = l
	logMu.Unlock()
}

// GetLogger get logger
func GetLogger() *log.Logger {
	return logger
}

// Logf is interface for logger
func Logf(ctx context.Context, isDebug bool, format string, v ...interface{}) {
	logMu.Lock()
	defer logMu.Unlock()

	prefix := GetLogPrefix(ctx, isDebug)
	format = prefix + format
	logger.Printf(format, v...)
}
