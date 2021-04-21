package models

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// PacketCarSetupData details the car setups for each vehicle in the session.
// Note that in multiplayer games, other player cars will appear as blank, you will only be able to see your car setup and AI cars.
type CarSetupData struct {
	Header    Header
	Timestamp time.Time

	CarSetups packet.CarSetupData
}

func NewCarSetupData(p *packet.PacketCarSetupData) *CarSetupData {
	return &CarSetupData{
		Header:    NewHeader(p.Header),
		Timestamp: time.Now().UTC(),
		CarSetups: p.CarSetups[p.Header.PlayerCarIndex],
	}
}

func (p *CarSetupData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
