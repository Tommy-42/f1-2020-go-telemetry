package models

import (
	"bytes"
	"encoding/json"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// FinalClassificationData details the final classification at the end of the race
// and the data will match with the post race results screen.
// This is especially useful for multiplayer games where it is not always possible to send lap times on the final frame because of network delay.
type FinalClassificationData struct {
	Header Header

	NumCars            uint8 // Number of cars in the final classification
	ClassificationData packet.FinalClassificationData
}

func NewFinalClassificationData(header packet.PacketHeader, p *packet.PacketFinalClassificationData) *FinalClassificationData {
	return &FinalClassificationData{
		Header:             NewHeader(header),
		NumCars:            p.NumCars,
		ClassificationData: p.ClassificationData[header.PlayerCarIndex],
	}
}

func (p *FinalClassificationData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
