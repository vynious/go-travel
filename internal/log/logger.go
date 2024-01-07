package log

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

var Global Logger

func init() {
	// Initialize the global logger instance here if needed
	Global = Logger{}
}

func getCallerInfo() (string, string) {
	pc, file, _, ok := runtime.Caller(2) // 2 levels up the stack
	if !ok {
		return "unknown", "unknown"
	}

	funcName := runtime.FuncForPC(pc).Name()
	shortFile := filepath.Base(file) // just the filename

	return shortFile, funcName
}

func (l Logger) log(level Level, args ...interface{}) {
	file, funcName := getCallerInfo()
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprint(args...)

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

func (l Logger) Info(args ...interface{}) {
	l.log(InfoLevel, args...)
}

func (l Logger) Warn(args ...interface{}) {
	l.log(WarnLevel, args...)
}

func (l Logger) Error(args ...interface{}) {
	l.log(ErrorLevel, args...)
}
