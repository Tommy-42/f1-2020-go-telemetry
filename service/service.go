package service

import (
	"context"
	"net"
	"os"

	"github.com/Tommy-42/f1-2020-go-telemetry/config"
	"github.com/Tommy-42/f1-2020-go-telemetry/repository/elastic"
	"github.com/Tommy-42/f1-2020-go-telemetry/service/handler"
	"github.com/sirupsen/logrus"
)

// Service ...
type Service struct {
	buffer []byte
}

func NewService() *Service {
	return &Service{
		buffer: make([]byte, 2048),
	}
}

// Start ...
func (s *Service) Start(ctx context.Context) error {

	config := config.Config{
		Elastic: elastic.DefaultConfig(),
	}

	if es := os.Getenv("ELASTICSEARCH_HOST"); es != "" {
		config.Elastic.Addresses = []string{es}
	}

	logrus.Info("starting New Repository")
	esRepo, err := elastic.NewES(config.Elastic)
	if err != nil {
		logrus.WithError(err).Error("could not start elastic repository")
		return err
	}

	logrus.Info("starting New Handler Packet")
	handlerPacket := handler.NewHandlerPacket(esRepo)

	logrus.Info("starting Listening on UDP port 20777")
	udp, err := net.ResolveUDPAddr("udp4", ":20777")
	if err != nil {
		logrus.WithError(err).Errorf("could not resolve udp addr with the port %s", "20777")
		return err
	}

	connection, err := net.ListenUDP("udp4", udp)
	if err != nil {
		logrus.WithError(err).Errorf("could not listening to udp socket")
		return err
	}
	defer connection.Close()

	buffer := make([]byte, 2048)
	for {
		if err := ctx.Err(); err != nil {
			logrus.WithError(err).Errorf("context cancelled")
			return err
		}

		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			logrus.WithError(err).Errorf("error reading from udp, received %d from %s:%d", n, addr.IP.String(), addr.Port)
		}
		// logrus.Debugf("reading from udp, received %d from %s:%d", n, addr.IP.String(), addr.Port)
		handlerPacket.HandlerChan() <- buffer
	}
}

func (s *Service) Stop() {}
