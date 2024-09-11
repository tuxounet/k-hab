package logger

import (
	"context"
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/tuxounet/k-hab/bases"
)

type Logger struct {
	name string
	log  *logrus.Entry
}

func NewLogger(ctx context.Context, name string, workFolder string) *Logger {
	rootLogger := logrus.New()
	rootLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		ForceColors:   true,
		FullTimestamp: true,
	})
	rootLogger.SetLevel(logrus.TraceLevel)

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile(path.Join(workFolder, "hab.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		rootLogger.Out = io.MultiWriter(os.Stdout, file)
	} else {
		rootLogger.Out = os.Stderr
		rootLogger.Warn("Failed to log to file, using default stderr")
	}
	return &Logger{
		name: name,
		log:  rootLogger.WithContext(ctx),
	}
}

func (l *Logger) GetName() string {
	return l.name
}
func (l *Logger) SetLevel(level logrus.Level) {
	l.log.Logger.SetLevel(level)
}

func (l *Logger) CreateSubLogger(name string, parent bases.ILogger) bases.ILogger {

	if parent != nil {
		name = parent.GetName() + "." + name
	}

	return &Logger{
		name: name,
		log:  l.log,
	}
}

func (l *Logger) TraceF(format string, args ...interface{}) {
	args = append([]interface{}{l.name}, args...)
	l.log.Tracef("%s > "+format, args...)
}
func (l *Logger) DebugF(format string, args ...interface{}) {

	args = append([]interface{}{l.name}, args...)
	l.log.Debugf("%s > "+format, args...)
}

func (l *Logger) InfoF(format string, args ...interface{}) {
	args = append([]interface{}{l.name}, args...)
	l.log.Infof("%s > "+format, args...)

}

func (l *Logger) WarnF(format string, args ...interface{}) {
	args = append([]interface{}{l.name}, args...)
	l.log.Warnf("%s > "+format, args...)

}

func (l *Logger) ErrorF(format string, args ...interface{}) {
	args = append([]interface{}{l.name}, args...)
	l.log.Errorf("%s > "+format, args...)
}
func (l *Logger) PanicF(format string, args ...interface{}) {
	args = append([]interface{}{l.name}, args...)
	l.log.Panicf("%s > "+format, args...)
}
