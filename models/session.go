package models

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// SessionData includes details about the current session in progress.
type SessionData struct {
	Header    Header
	Timestamp time.Time

	// Weather - 0 = clear, 1 = light cloud, 2 = overcast
	// 3 = light rain, 4 = heavy rain, 5 = storm
	Weather uint8

	TrackTemperature int8   // Track temp. in degrees celsius
	AirTemperature   int8   // Air temp. in degrees celsius
	TotalLaps        uint8  // Total number of laps in this race
	TrackLength      uint16 // Track length in metres
	// 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P
	// 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ
	// 10 = R, 11 = R2, 12 = Time Trial
	SessionType uint8
	TrackID     int8 // -1 for unknown, 0-21 for tracks, see appendix
	// Formula
	// 0 = F1 Modern
	// 1 = F1 Classic
	// 2 = F2
	// 3 = F1 Generic
	Formula             uint8
	SessionTimeLeft     uint16             // Time left in session in seconds
	SessionDuration     uint16             // Session duration in seconds
	PitSpeedLimit       uint8              // Pit speed limit in kilometres per hour
	GamePaused          uint8              // Whether the game is paused
	IsSpectating        uint8              // Whether the player is spectating
	SpectatorCarIndex   uint8              // Index of the car being spectated
	SliProNativeSupport uint8              // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones     uint8              // Number of marshal zones to follow
	MarshalZones        packet.MarshalZone // List of marshal zones – max 21
	// SafetyCarStatus
	// 0 = no safety car
	// 1 = full safety car
	// 2 = virtual safety car
	SafetyCarStatus           uint8
	NetworkGame               uint8                        // 0 = offline, 1 = online
	NumWeatherForecastSamples uint8                        // Number of weather samples to follow
	WeatherForecastSamples    packet.WeatherForecastSample // Array of weather forecast samples
}

func NewSessionData(p *packet.PacketSessionData) *SessionData {
	return &SessionData{
		Header:    NewHeader(p.Header),
		Timestamp: time.Now().UTC(),

		Weather:                   p.Weather,
		TrackTemperature:          p.TrackTemperature,
		AirTemperature:            p.AirTemperature,
		TotalLaps:                 p.TotalLaps,
		TrackLength:               p.TrackLength,
		SessionType:               p.SessionType,
		TrackID:                   p.TrackID,
		Formula:                   p.Formula,
		SessionTimeLeft:           p.SessionTimeLeft,
		SessionDuration:           p.SessionDuration,
		PitSpeedLimit:             p.PitSpeedLimit,
		GamePaused:                p.GamePaused,
		IsSpectating:              p.IsSpectating,
		SpectatorCarIndex:         p.SpectatorCarIndex,
		SliProNativeSupport:       p.SliProNativeSupport,
		NumMarshalZones:           p.NumMarshalZones,
		MarshalZones:              p.MarshalZones[0], // FIXME: understand whats needed here
		SafetyCarStatus:           p.SafetyCarStatus,
		NetworkGame:               p.NetworkGame,
		NumWeatherForecastSamples: p.NumWeatherForecastSamples,
		WeatherForecastSamples:    p.WeatherForecastSamples[0],
	}
}

func (p *SessionData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
