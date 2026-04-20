// Full Go Log System with Rotation

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// ============================
// LOGGER STRUCT
// ============================
type Logger struct {
	file       *os.File
	filename   string
	maxSize    int64
	mu         sync.Mutex
}

// ============================
// CREATE NEW LOGGER
// ============================
func NewLogger(filename string, maxSize int64) *Logger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return &Logger{
		file:     file,
		filename: filename,
		maxSize:  maxSize,
	}
}

// ============================
// CHECK ROTATION NEED
// ============================
func (l *Logger) rotateIfNeeded() {
	info, err := l.file.Stat()
	if err != nil {
		return
	}

	if info.Size() < l.maxSize {
		return
	}

	l.file.Close()

	// rename old logs
	timestamp := time.Now().Format("20060102_150405")
	newName := fmt.Sprintf("%s.%s", l.filename, timestamp)

	os.Rename(l.filename, newName)

	// create new file
	file, err := os.OpenFile(l.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	l.file = file
}

// ============================
// WRITE LOG
// ============================
func (l *Logger) Log(message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.rotateIfNeeded()

	logLine := fmt.Sprintf("[%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		message,
	)

	l.file.WriteString(logLine)
}

// ============================
// INFO / ERROR / DEBUG HELPERS
// ============================
func (l *Logger) Info(msg string) {
	l.Log("INFO: " + msg)
}

func (l *Logger) Error(msg string) {
	l.Log("ERROR: " + msg)
}

func (l *Logger) Debug(msg string) {
	l.Log("DEBUG: " + msg)
}

// ============================
// CLOSE LOGGER
// ============================
func (l *Logger) Close() {
	l.file.Close()
}

// ============================
// DEMO APP
// ============================
func main() {
	logger := NewLogger("app.log", 1*1024) // 1 KB max size for demo
	defer logger.Close()

	fmt.Println("Logging system started...")

	for i := 0; i < 100; i++ {
		logger.Info(fmt.Sprintf("Processing request #%d", i))
		logger.Debug("Debugging system flow")
		logger.Error("Sample error message")

		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("Done logging. Check rotated files.")
}