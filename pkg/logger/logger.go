package logger

import (
    "log"
    "os"
)

type Logger struct {
    infoLogger  *log.Logger
    errorLogger *log.Logger
}

func New() *Logger {
    return &Logger{
        infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
        errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
    }
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
    l.infoLogger.Println(msg, keysAndValues)
}

func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
    l.errorLogger.Println(msg, keysAndValues)
}