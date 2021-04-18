package handler

import (
	"bytes"
	"context"
	"encoding/binary"

	"github.com/mailgun/holster/v3/syncutil"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Tommy-42/f1-2020-go-telemetry/models"
	f1packet "github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
	"github.com/Tommy-42/f1-2020-go-telemetry/repository"
)

const (
	F1Version uint16 = 2020
)

// HandlerPacket ...
type HandlerPacket struct {
	repo repository.Repository

	handlerChan chan []byte
}

// NewHandlerPacket ...
func NewHandlerPacket(repo repository.Repository) *HandlerPacket {
	h := &HandlerPacket{
		repo:        repo,
		handlerChan: make(chan []byte, 10000),
	}

	ctx := context.Background()
	go func() {

		fanOut := syncutil.NewFanOut(100)

		for pkt := range h.handlerChan {
			fanOut.Run(func(cast interface{}) error {
				// avoid data race
				packet := cast.([]byte)

				// decode packet to the correct one
				data, err := h.decodePacket(ctx, packet)
				if err != nil {
					if errors.Is(err, ErrIgnorePacket) {
						return nil
					}
					logrus.WithError(err).Errorf("found error while decoding packet")
					return err

				}

				// handle packet
				err = h.storeData(ctx, data)
				if err != nil {
					logrus.WithError(err).Errorf("found error while handling packet")
				}
				return err
			}, pkt)
		}
	}()

	return h
}

// HandlerChan ...
func (h *HandlerPacket) HandlerChan() chan []byte {
	return h.handlerChan
}

// decodePacket ...
func (h *HandlerPacket) decodePacket(ctx context.Context, packet []byte) (models.F1Data, error) {

	headerReader := bytes.NewReader(packet)
	header := f1packet.PacketHeader{}
	err := binary.Read(headerReader, binary.LittleEndian, &header)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode header")
	}

	if header.PacketFormat != F1Version || header.PlayerCarIndex > 21 {
		return nil, ErrIgnorePacket
	}

	switch f1packet.PacketType(header.PacketID) {
	case f1packet.MotionPacket:

		placeholder := &f1packet.PacketMotionData{}
		reader := bytes.NewReader(packet)
		err = binary.Read(reader, binary.LittleEndian, placeholder)
		if err != nil {
			logrus.Errorf("Packet ID: %d", header.PacketID)
			return nil, errors.Wrap(err, "could not decode binary data PacketMotionData")
		}

		data := models.NewMotionData(placeholder)
		if data == nil {
			return nil, ErrIgnorePacket
		}
		return data, nil

	case f1packet.SessionPacket:

		placeholder := &f1packet.PacketSessionData{}
		reader := bytes.NewReader(packet)
		err = binary.Read(reader, binary.LittleEndian, placeholder)
		if err != nil {
			logrus.Errorf("Packet ID: %d", header.PacketID)
			return nil, errors.Wrap(err, "could not decode binary data PacketSessionData")
		}

		data := models.NewSessionData(placeholder)
		if data == nil {
			return nil, ErrIgnorePacket
		}
		return data, nil

	case f1packet.LapDataPacket:

		placeholder := &f1packet.PacketLapData{}
		reader := bytes.NewReader(packet)
		err = binary.Read(reader, binary.LittleEndian, placeholder)
		if err != nil {
			logrus.Errorf("Packet ID: %d", header.PacketID)
			return nil, errors.Wrap(err, "could not decode binary data PacketLapData")
		}

		data := models.NewLapData(placeholder)
		if data == nil {
			return nil, ErrIgnorePacket
		}
		return data, nil

	case f1packet.EventPacket:

		placeholder := &f1packet.PacketEventData{}
		reader := bytes.NewReader(packet)
		err = binary.Read(reader, binary.LittleEndian, placeholder)
		if err != nil {
			logrus.Errorf("Packet ID: %d", header.PacketID)
			return nil, errors.Wrap(err, "could not decode binary data PacketEventData")
		}

		data := models.NewEventData(placeholder)
		if data == nil {
			return nil, ErrIgnorePacket
		}
		return data, nil

	case f1packet.ParticipantsPacket:

		placeholder := &f1packet.PacketParticipantsData{}
		reader := bytes.NewReader(packet)
		err = binary.Read(reader, binary.LittleEndian, placeholder)
		if err != nil {
			logrus.Errorf("Packet ID: %d", header.PacketID)
			return nil, errors.Wrap(err, "could not decode binary data PacketParticipantsData")
		}

		data := models.NewParticipantsData(placeholder)
		if data == nil {
			return nil, ErrIgnorePacket
		}
		return data, nil

	case f1packet.CarSetupsPacket:

		placeholder := &f1packet.PacketCarSetupData{}
		reader := bytes.NewReader(packet)
		err = binary.Read(reader, binary.LittleEndian, placeholder)
		if err != nil {
			logrus.Errorf("Packet ID: %d", header.PacketID)
			return nil, errors.Wrap(err, "could not decode binary data PacketCarSetupData")
		}

		data := models.NewCarSetupData(placeholder)
		if data == nil {
			return nil, ErrIgnorePacket
		}
		return data, nil

	case f1packet.CarTelemetryPacket:

		placeholder := &f1packet.PacketCarTelemetryData{}
		reader := bytes.NewReader(packet)
		err = binary.Read(reader, binary.LittleEndian, placeholder)
		if err != nil {
			logrus.Errorf("Packet ID: %d", header.PacketID)
			return nil, errors.Wrap(err, "could not decode binary data PacketCarTelemetryData")
		}

		data := models.NewCarTelemetryData(placeholder)
		if data == nil {
			return nil, ErrIgnorePacket
		}
		return data, nil

	case f1packet.CarStatusPacket:

		placeholder := &f1packet.PacketCarStatusData{}
		reader := bytes.NewReader(packet)
		err = binary.Read(reader, binary.LittleEndian, placeholder)
		if err != nil {
			logrus.Errorf("Packet ID: %d", header.PacketID)
			return nil, errors.Wrap(err, "could not decode binary data PacketCarStatusData")
		}

		data := models.NewCarStatusData(placeholder)
		if data == nil {
			return nil, ErrIgnorePacket
		}
		return data, nil

	case f1packet.FinalClassificationPacket:

		placeholder := &f1packet.PacketFinalClassificationData{}
		reader := bytes.NewReader(packet)
		err = binary.Read(reader, binary.LittleEndian, placeholder)
		if err != nil {
			logrus.Errorf("Packet ID: %d", header.PacketID)
			return nil, errors.Wrap(err, "could not decode binary data PacketFinalClassificationData")
		}

		data := models.NewFinalClassificationData(placeholder)
		if data == nil {
			return nil, ErrIgnorePacket
		}
		return data, nil

	case f1packet.LobbyInfoPacket:

		placeholder := &f1packet.PacketLobbyInfoData{}
		reader := bytes.NewReader(packet)
		err = binary.Read(reader, binary.LittleEndian, placeholder)
		if err != nil {
			logrus.Errorf("Packet ID: %d", header.PacketID)
			return nil, errors.Wrap(err, "could not decode binary data PacketLobbyInfoData")
		}

		data := models.NewLobbyInfoData(placeholder)
		if data == nil {
			return nil, ErrIgnorePacket
		}
		return data, nil
	default:
		return nil, ErrUnknownPacket
	}
}

func (h *HandlerPacket) storeData(ctx context.Context, data models.F1Data) error {

	body, err := data.ToJson()
	if err != nil {
		return errors.Wrap(err, "could not convert data to json")
	}

	err = h.repo.Store(ctx, body)
	if err != nil {
		return errors.Wrap(err, "could not handle packet")
	}

	return nil
}
