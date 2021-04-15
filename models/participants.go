package models

import (
	"bytes"
	"encoding/json"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// ParticipantsData is a list of participants in the race.
// If the vehicle is controlled by AI, then the name will be the driver name.
// If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.
//
// N.B. on Xbox One, the names will always be the driver name, on PS4 the name will be the LAN name if playing a LAN game, otherwise it will be the driver name.
//
// The array should be indexed by vehicle index.
//
// Frequency: Every 5 seconds
// Size: 1213 bytes (Packet size updated in Beta 3)
// Version: 1
type ParticipantsData struct {
	Header Header
	// Number of active cars in the data – should match number of cars on HUD
	NumActiveCars uint8
	Participants  packet.ParticipantData
}

func NewParticipantsData(p *packet.PacketParticipantsData) *ParticipantsData {
	if p.Header.PlayerCarIndex > 21 {
		return nil
	}

	return &ParticipantsData{
		Header:        NewHeader(p.Header),
		NumActiveCars: p.NumActiveCars,
		Participants:  p.Participants[p.Header.PlayerCarIndex],
	}
}

func (p *ParticipantsData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
