package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Entry
}

func NewLogger(ctx context.Context) *Logger {
	rootLogger := logrus.New()
	rootLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "15-01-2018 15:04:05.000000",
	})
	rootLogger.SetLevel(logrus.TraceLevel)
	return &Logger{
		log: rootLogger.WithContext(ctx),
	}
}

func (l *Logger) SetLevel(level logrus.Level) {
	l.log.Logger.SetLevel(level)
}

func (l *Logger) CreateScopeLogger(name string, fields map[string]interface{}) *Logger {
	return &Logger{
		log: l.log.WithField("scope", name).WithFields(fields),
	}
}

func (l *Logger) TraceF(format string, args ...interface{}) {
	l.log.Tracef(format, args...)
}
func (l *Logger) DebugF(format string, args ...interface{}) {

	l.log.Debugf(format, args...)
}

func (l *Logger) InfoF(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *Logger) WarnF(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *Logger) PanicF(format string, args ...interface{}) {
	l.log.Panicf(format, args...)
}
