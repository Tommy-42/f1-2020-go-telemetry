package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tommy-42/f1-go-telemetry/service"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	svc := service.NewService()

	err := svc.Start()
	if err != nil {
		logrus.WithError(err).Errorf("error starting service")
	}
}
