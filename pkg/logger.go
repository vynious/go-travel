package pkg

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/fatih/color"
)

type Level string

const (
	InfoLevel  Level = "INFO"
	WarnLevel  Level = "WARN"
	ErrorLevel Level = "ERROR"
)

type Logger struct{}

var Log Logger

func init() {
	// Initialize the global logger instance here if needed
	Log = Logger{}
}

func getCallerInfo() (string, string) {
	// Ascend 3 levels up the stack:
	// 1. getCallerInfo itself
	// 2. the log method (Info, Warn, Error, etc.)
	// 3. the actual caller of the log method
	pc, file, _, ok := runtime.Caller(3) // Change this to 3
	if !ok {
		return "unknown", "unknown"
	}

	funcName := runtime.FuncForPC(pc).Name()
	shortFile := filepath.Base(file) // just the filename

	return shortFile, funcName
}

func (l Logger) log(level Level, format string, args ...interface{}) {
	file, funcName := getCallerInfo()
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprintf(format, args...)

	switch level {
	case InfoLevel:
		color.New(color.FgBlue).Printf("[%s] [%s] [%s -> %s] %s\n", timestamp, level, file, funcName, message)
	case WarnLevel:
		color.New(color.FgYellow).Printf("[%s] [%s] [%s -> %s] %s\n", timestamp, level, file, funcName, message)
	case ErrorLevel:
		color.New(color.FgRed).Printf("[%s] [%s] [%s -> %s] %s\n", timestamp, level, file, funcName, message)
	default:
		fmt.Printf("[%s] [%s] [%s -> %s] %s\n", timestamp, level, file, funcName, message)
	}
}

func (l Logger) Info(format string, args ...interface{}) {
	l.log(InfoLevel, format, args...)
}

func (l Logger) Warn(format string, args ...interface{}) {
	l.log(WarnLevel, format, args...)
}

func (l Logger) Error(format string, args ...interface{}) {
	l.log(ErrorLevel, format, args...)
}
