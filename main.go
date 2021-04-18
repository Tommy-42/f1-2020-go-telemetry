package main

import (
	"context"
	"os"

	"github.com/Tommy-42/f1-2020-go-telemetry/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	svc := service.NewService()

	err := svc.Start(context.Background())
	if err != nil {
		logrus.WithError(err).Errorf("error starting service")
		os.Exit(1)
	}

	// handle signal

	svc.Stop()

	os.Exit(0)
}
