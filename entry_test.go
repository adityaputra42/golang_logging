package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithFields(logrus.Fields{
		"name":  "Paijo",
		"email": "paijo98@gmail.com",
	})
	entry.Info("Hello Entry!")

}
