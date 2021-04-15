package models

import (
	"bytes"
	"encoding/json"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// CarStatusData details car statuses for all the cars in the race.
// It includes values such as the damage readings on the car.
type CarStatusData struct {
	Header Header

	CarStatusData packet.CarStatusData
}

func NewCarStatusData(p *packet.PacketCarStatusData) *CarStatusData {
	if p.Header.PlayerCarIndex > 21 {
		return nil
	}
	return &CarStatusData{
		Header:        NewHeader(p.Header),
		CarStatusData: p.CarStatusData[p.Header.PlayerCarIndex],
	}
}

func (p *CarStatusData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
