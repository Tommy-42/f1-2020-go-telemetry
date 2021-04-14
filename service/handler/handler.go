package handler

import (
	"github.com/Tommy-42/f1-2020-go-telemetry/models"
	"github.com/mailgun/holster/v3/syncutil"
)

// HandlerPacket ...
type HandlerPacket struct {
	handlerChan chan []byte
}

// NewHandlerPacket ...
func NewHandlerPacket() *HandlerPacket {
	return &HandlerPacket{
		handlerChan: make(chan []byte, 100),
	}
}

// HandlerChan ...
func (h *HandlerPacket) HandlerChan() chan []byte {
	return h.handlerChan
}

// Run ...
func (h *HandlerPacket) Run() {

	fanOut := syncutil.NewFanOut(100)
	for packet := range h.handlerChan {
		fanOut.Run(func(cast interface{}) error {
			// find
			p := h.findPacketType(packet)

			// handle packet
			err := h.handlePacket(p)
			return err
		}, packet)
	}
}

// findPacketType ...
func (h *HandlerPacket) findPacketType(models.X) error {

	return nil
}

// HandlePacket ...
func (h *HandlerPacket) handlePacket(models.X) error {

	return nil
}
