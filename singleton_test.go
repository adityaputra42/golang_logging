package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {

	logrus.Info("Hallo Singleton info")
	logrus.Warn("Hallo Singleton warning")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hallo Singleton info")
	logrus.Warn("Hallo Singleton warning")

}
