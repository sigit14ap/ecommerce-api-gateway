package logger

import (
	"io"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))

	logFile := &lumberjack.Logger{
		Filename:   "logs/api-gateway.log",
		MaxSize:    10,
		MaxBackups: 30,
		MaxAge:     1,
		Compress:   false,
	}

	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	log.SetFormatter(&logrus.JSONFormatter{})

	switch logLevel {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
}

func Info(msg string, fields logrus.Fields) {
	log.WithFields(fields).Info(msg)
}

func Error(msg string, fields logrus.Fields) {
	log.WithFields(fields).Error(msg)
}

func Warn(msg string, fields logrus.Fields) {
	log.WithFields(fields).Warn(msg)
}
