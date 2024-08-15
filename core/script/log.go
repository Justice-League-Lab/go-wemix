package script

import (
	"fmt"
	"io"

	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger = &lumberjack.Logger{
	Filename:   "",
	MaxSize:    10, // megabytes
	MaxBackups: 300,
	MaxAge:     60, //days
}

type LoggerLevel string

const (
	LogLevelPanic  LoggerLevel = "panic"
	LogLevelFatal  LoggerLevel = "fatal"
	LogLevelError  LoggerLevel = "error"
	LogLevelWarn   LoggerLevel = "warn"
	LogLevelInfo   LoggerLevel = "info"
	LogLevelDebug  LoggerLevel = "debug"
	LogLevelTrance LoggerLevel = "trance"
)

func InitLog(logFilePath, logName, logLevel string) {
	logFileName := filepath.Join(logFilePath, logName)
	logger.Filename = logFileName
	//logrus.JSONFormatter{} and logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	fileAndStdoutWriter := io.Writer(logger)
	logrus.SetOutput(fileAndStdoutWriter)
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		fmt.Printf("parse logs level err, %v, use defaule level 'info'", err)
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)

	logrus.Info("current log level is " + logLevel)
}
