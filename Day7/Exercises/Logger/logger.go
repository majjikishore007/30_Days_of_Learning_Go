package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	writer io.Writer
}

func NewLogger(writer io.Writer) *Logger {
	return &Logger{writer: writer}
}

func (l *Logger) getTimeStamp() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

func (l *Logger) writeLog(level, txt string) {
	fmt.Fprintf(l.writer, "%s [%s] %s\n", l.getTimeStamp(), level, txt)
}

func (l *Logger) Error(txt string) {
	l.writeLog("ERROR", txt)
}
func (l *Logger) Info(txt string) {
	l.writeLog("INFO", txt)

}
func (l *Logger) Debug(txt string) {
	l.writeLog("DEBUG", txt)

}

func main() {

	// single writer
	// logger.Debug("hey there")
	// logger.Error("index out of bound")
	// logger.Info("This is just to test")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileLogger := NewLogger(file)
	fileLogger.Error("Something went wrong")

	// Bonus: Log to multiple destinations
	multiLogger := NewLogger(io.MultiWriter(os.Stdout, file))
	multiLogger.Debug("This goes to both stdout and file")

}
