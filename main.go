package main

import (
	"github.com/Tommy-42/f1-2020-go-telemetry/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	svc := service.NewService()

	err := svc.Start()
	if err != nil {
		logrus.WithError(err).Errorf("error starting service")
	}
}
