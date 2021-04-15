package models

import (
	"bytes"
	"encoding/json"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// CarTelemetryData details telemetry for all the cars in the race.
// It details various values that would be recorded on the car such as speed, throttle application, DRS etc.
type CarTelemetryData struct {
	Header Header

	CarTelemetryData packet.CarTelemetryData

	// Bit flags specifying which buttons are being pressed
	// currently - see appendices
	ButtonStatus uint32

	// Added in Beta 3:
	// Index of MFD panel open - 255 = MFD closed
	// Single player, race – 0 = Car setup, 1 = Pits
	// 2 = Damage, 3 =  Engine, 4 = Temperatures
	// May vary depending on game mode
	MfdPanelIndex                uint8
	MfdPanelIndexSecondaryPlayer uint8 // See above

	// Suggested gear for the player (1-8)
	// 0 if no gear suggested
	SuggestedGear int8
}

func NewCarTelemetryData(p *packet.PacketCarTelemetryData) *CarTelemetryData {
	if p.Header.PlayerCarIndex > 21 {
		return nil
	}

	return &CarTelemetryData{
		Header:                       NewHeader(p.Header),
		CarTelemetryData:             p.CarTelemetryData[p.Header.PlayerCarIndex],
		ButtonStatus:                 p.ButtonStatus,
		MfdPanelIndex:                p.MfdPanelIndex,
		MfdPanelIndexSecondaryPlayer: p.MfdPanelIndexSecondaryPlayer,
		SuggestedGear:                p.SuggestedGear,
	}
}

func (p *CarTelemetryData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
