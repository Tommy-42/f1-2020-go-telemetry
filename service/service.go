package service

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/tommy-42/f1-go-telemetry/handler"
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
func (s *Service) Start() error {

	handlerPacket := handler.NewHandlerPacket()

	udp, err := net.ResolveUDPAddr("udp4", ":20777")
	if err != nil {
		logrus.WithError(err).Errorf("could not resolve udp addr with the port %d", "20777")
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
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			logrus.WithError(err).Errorf("error reading from udp, received %d from %s:%s", n, addr.IP, addr.Port)
		}
		handlerPacket.HandlerChan <- buffer
	}

	return nil
}

func (s *Service) Stop() {
	return
}
