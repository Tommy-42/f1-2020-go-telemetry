package models

import (
	"bytes"
	"encoding/json"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

type CarTelemetryDataDetails struct {
	Speed            uint16  // Speed of car in kilometres per hour
	Throttle         float64 // Amount of throttle applied (0.0 to 1.0)
	Steer            float64 // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake            float64 // Amount of brake applied (0.0 to 1.0)
	Clutch           uint8   // Amount of clutch applied (0 to 100)
	Gear             int8    // Gear selected (1-8, N=0, R=-1)
	EngineRPM        uint16  // Engine RPM
	Drs              uint8   // 0 = off, 1 = on
	RevLightsPercent uint8   // Rev lights indicator (percentage)

	RearLeftBrakesTemperature   uint16 // Brakes temperature (celsius)
	RearRightBrakesTemperature  uint16 // Brakes temperature (celsius)
	FrontLeftBrakesTemperature  uint16 // Brakes temperature (celsius)
	FrontRightBrakesTemperature uint16 // Brakes temperature (celsius)

	RearLeftTyresSurfaceTemperature   uint8 // Tyres surface temperature (celsius)
	RearRightTyresSurfaceTemperature  uint8 // Tyres surface temperature (celsius)
	FrontLeftTyresSurfaceTemperature  uint8 // Tyres surface temperature (celsius)
	FrontRightTyresSurfaceTemperature uint8 // Tyres surface temperature (celsius)

	RearLeftTyresInnerTemperature   uint8 // Tyres inner temperature (celsius)
	RearRightTyresInnerTemperature  uint8 // Tyres inner temperature (celsius)
	FrontLeftTyresInnerTemperature  uint8 // Tyres inner temperature (celsius)
	FrontRightTyresInnerTemperature uint8 // Tyres inner temperature (celsius)

	EngineTemperature uint16 // Engine temperature (celsius)

	RearLeftTyresPressure   float64 // Tyres pressure (PSI)
	RearRightTyresPressure  float64 // Tyres pressure (PSI)
	FrontLeftTyresPressure  float64 // Tyres pressure (PSI)
	FrontRightTyresPressure float64 // Tyres pressure (PSI)

	RearLeftSurfaceType   uint8 // Driving surface, see appendices
	RearRightSurfaceType  uint8 // Driving surface, see appendices
	FrontLeftSurfaceType  uint8 // Driving surface, see appendices
	FrontRightSurfaceType uint8 // Driving surface, see appendices
}

// CarTelemetryData details telemetry for all the cars in the race.
// It details various values that would be recorded on the car such as speed, throttle application, DRS etc.
type CarTelemetryData struct {
	Header Header

	CarTelemetryData CarTelemetryDataDetails

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

func NewCarTelemetryData(header packet.PacketHeader, p *packet.PacketCarTelemetryData) *CarTelemetryData {
	pk := p.CarTelemetryData[header.PlayerCarIndex]
	return &CarTelemetryData{
		Header: NewHeader(header),
		CarTelemetryData: CarTelemetryDataDetails{
			Speed:                             pk.Speed,
			Throttle:                          pk.Throttle,
			Steer:                             pk.Steer,
			Brake:                             pk.Brake,
			Clutch:                            pk.Clutch,
			Gear:                              pk.Gear,
			EngineRPM:                         pk.EngineRPM,
			Drs:                               pk.Drs,
			RevLightsPercent:                  pk.RevLightsPercent,
			RearLeftBrakesTemperature:         pk.BrakesTemperature[0],
			RearRightBrakesTemperature:        pk.BrakesTemperature[1],
			FrontLeftBrakesTemperature:        pk.BrakesTemperature[2],
			FrontRightBrakesTemperature:       pk.BrakesTemperature[3],
			RearLeftTyresSurfaceTemperature:   pk.TyresSurfaceTemperature[0],
			RearRightTyresSurfaceTemperature:  pk.TyresSurfaceTemperature[1],
			FrontLeftTyresSurfaceTemperature:  pk.TyresSurfaceTemperature[2],
			FrontRightTyresSurfaceTemperature: pk.TyresSurfaceTemperature[3],
			RearLeftTyresInnerTemperature:     pk.TyresInnerTemperature[0],
			RearRightTyresInnerTemperature:    pk.TyresInnerTemperature[1],
			FrontLeftTyresInnerTemperature:    pk.TyresInnerTemperature[2],
			FrontRightTyresInnerTemperature:   pk.TyresInnerTemperature[3],
			EngineTemperature:                 pk.EngineTemperature,
			RearLeftTyresPressure:             pk.TyresPressure[0],
			RearRightTyresPressure:            pk.TyresPressure[1],
			FrontLeftTyresPressure:            pk.TyresPressure[2],
			FrontRightTyresPressure:           pk.TyresPressure[3],
			RearLeftSurfaceType:               pk.SurfaceType[0],
			RearRightSurfaceType:              pk.SurfaceType[1],
			FrontLeftSurfaceType:              pk.SurfaceType[2],
			FrontRightSurfaceType:             pk.SurfaceType[3],
		},
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
