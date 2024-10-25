package utils

import (
	"log"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

// if skip == 0; return none field logger
// else; return be called file and function field
func NewLogger(skip int) *logrus.Entry {
	if skip == 0 {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logger := logrus.WithFields(logrus.Fields{})
		return logger
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		pc, file, _, _ := runtime.Caller(skip)
		logger := logrus.WithFields(logrus.Fields{
			"file":     file,
			"function": runtime.FuncForPC(pc).Name(),
		})
		return logger
	}
}

func ErrLogging() {
	// [local/fin/utils.b]: 15Line
	log_format := "[%v]: %vL"
	pc := make([]uintptr, 10)
	n := runtime.Callers(2, pc)

	pc = pc[:n]
	frames := runtime.CallersFrames(pc)
	for i := 0; i < n; i++ {
		frame, more := frames.Next()
		if !strings.Contains(frame.Function, "local/fin") {
			continue
		} else if !more {
			break
		} else {
			log.Printf(log_format, frame.Function, frame.Line)
		}
	}
}
