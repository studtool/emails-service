package logs

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: func() *logrus.Logger {
			log := logrus.StandardLogger()
			log.SetFormatter(&logrus.JSONFormatter{})
			return log
		}(),
	}
}

func (log *Logger) Debug(args ...interface{}) {
	log.logger.Debug(args...)
}

func (log *Logger) Info(args ...interface{}) {
	log.logger.Info(args...)
}

func (log *Logger) Warning(args ...interface{}) {
	log.logger.Warn(args...)
}

func (log *Logger) Error(args ...interface{}) {
	log.logger.Error(args...)
}

func (log *Logger) Fatal(args ...interface{}) {
	log.logger.Fatal(args...)
}
