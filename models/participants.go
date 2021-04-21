package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// ParticipantData ...
type ParticipantData struct {
	AiControlled  uint8  // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID      uint8  // Driver id - see appendix
	TeamID        uint8  // Team id - see appendix
	RaceNumber    uint8  // Race number of the car
	Nationality   uint8  // Nationality of the driver
	Name          string // Name of participant in UTF-8 format – null terminated. Will be truncated with … (U+2026) if too long
	YourTelemetry uint8  // The player's UDP setting, 0 = restricted, 1 = public
}

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
	Header    Header
	Timestamp time.Time

	// Number of active cars in the data – should match number of cars on HUD
	NumActiveCars uint8
	Participants  ParticipantData
}

func NewParticipantsData(p *packet.PacketParticipantsData) *ParticipantsData {
	pk := p.Participants[p.Header.PlayerCarIndex]
	return &ParticipantsData{
		Header:    NewHeader(p.Header),
		Timestamp: time.Now().UTC(),

		NumActiveCars: p.NumActiveCars,
		Participants: ParticipantData{
			AiControlled:  pk.AiControlled,
			DriverID:      pk.DriverID,
			TeamID:        pk.TeamID,
			RaceNumber:    pk.RaceNumber,
			Nationality:   pk.Nationality,
			Name:          fmt.Sprintf("%s", pk.Name),
			YourTelemetry: pk.YourTelemetry,
		},
	}
}

func (p *ParticipantsData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
