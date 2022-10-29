package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

// formatter can be json or text with default as "json", out can be anything that implements io.Writter interface
// with os.Stderr as default, logLevel specifies logger log level.
func New(formatter string, out io.Writer, logLevel int) *Logger {

	var log = logrus.New()

	if formatter != "json" {
		log.SetFormatter(&logrus.JSONFormatter{})
	}

	if logLevel != -1 {
		log.Level = logrus.Level(logLevel)
	}

	return &Logger{logger: log}
}

func (l Logger) Trace(args string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Trace(args)
}

func (l Logger) Debug(args string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Debug(args)
}

func (l Logger) Info(args string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Info(args)
}

func (l Logger) Warn(args string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Warn(args)
}

func (l Logger) Error(args string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Error(args)
}

func (l Logger) Fatal(msg string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Fatal(msg)
}

func (l Logger) Panic(args string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Panic(args)
}
