package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.WithField("name", "Aditya").Info("Hello Bro!")

}
func TestFields(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.WithFields(logrus.Fields{
		"name":  "Aditya",
		"email": "aditya27@gmail.com",
		"age":   27,
	}).Info("Hello Bro!")

}
