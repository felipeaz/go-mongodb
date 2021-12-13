package logrus

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewJsonLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	return log
}

func NewTextLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	return log
}
