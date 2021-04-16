package models

import (
	"bytes"
	"encoding/json"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// LapData contains the LapData for all the cars on track
type LapData struct {
	Header Header

	LapData packet.LapData // Lap data for all cars on track
}

func NewLapData(header packet.PacketHeader, p *packet.PacketLapData) *LapData {
	return &LapData{
		Header:  NewHeader(header),
		LapData: p.LapData[header.PlayerCarIndex],
	}
}

func (p *LapData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
