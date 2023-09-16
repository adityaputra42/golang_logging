package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormater(t *testing.T) {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})

	log.Printf("Halo format data json")

}
