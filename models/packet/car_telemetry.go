package packet

import (
	"encoding/json"
)

// CarTelemetryData ...
type CarTelemetryData struct {
	Speed                   uint16     // Speed of car in kilometres per hour
	Throttle                float64    // Amount of throttle applied (0.0 to 1.0)
	Steer                   float64    // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake                   float64    // Amount of brake applied (0.0 to 1.0)
	Clutch                  uint8      // Amount of clutch applied (0 to 100)
	Gear                    int8       // Gear selected (1-8, N=0, R=-1)
	EngineRPM               uint16     // Engine RPM
	Drs                     uint8      // 0 = off, 1 = on
	RevLightsPercent        uint8      // Rev lights indicator (percentage)
	BrakesTemperature       [4]uint16  // Brakes temperature (celsius)
	TyresSurfaceTemperature [4]uint8   // Tyres surface temperature (celsius)
	TyresInnerTemperature   [4]uint8   // Tyres inner temperature (celsius)
	EngineTemperature       uint16     // Engine temperature (celsius)
	TyresPressure           [4]float64 // Tyres pressure (PSI)
	SurfaceType             [4]uint8   // Driving surface, see appendices
}

// PacketCarTelemetryData details telemetry for all the cars in the race.
// It details various values that would be recorded on the car such as speed, throttle application, DRS etc.
//
// Frequency: Rate as specified in menus
// Size: 1307 bytes (Packet size updated in Beta 3)
// Version: 1
type PacketCarTelemetryData struct {
	Header PacketHeader // Header

	CarTelemetryData [22]CarTelemetryData

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

func (p *PacketCarTelemetryData) Read(receiver []byte) (n int, err error) {
	data, err := json.Marshal(p)
	copy(receiver, data)
	n = len(data)
	return
}
