package models

import (
	"bytes"
	"encoding/json"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// CarStatusData ...
type CarStatusDataDetail struct {
	TractionControl   uint8   // 0 (off) - 2 (high)
	AntiLockBrakes    uint8   // 0 (off) - 1 (on)
	FuelMix           uint8   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias    uint8   // Front brake bias (percentage)
	PitLimiterStatus  uint8   // Pit limiter status - 0 = off, 1 = on
	FuelInTank        float64 // Current fuel mass
	FuelCapacity      float64 // Fuel capacity
	FuelRemainingLaps float64 // Fuel remaining in terms of laps (value on MFD)
	MaxRPM            uint16  // Cars max RPM, point of rev limiter
	IdleRPM           uint16  // Cars idle RPM
	MaxGears          uint8   // Maximum number of gears
	DrsAllowed        uint8   // 0 = not allowed, 1 = allowed, -1 = unknown

	// Added in Beta3:
	// 0 = DRS not available, non-zero - DRS will be available
	// in [X] metres
	DrsActivationDistance uint16

	// Tyre wear percentage
	// F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1
	// 7 = inter, 8 = wet
	// F1 Classic - 9 = dry, 10 = wet
	// F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard
	// 15 = wet
	RearLeftTyresWear   uint8
	RearRightTyresWear  uint8
	FrontLeftTyresWear  uint8
	FrontRightTyresWear uint8

	ActualTyreCompound uint8
	// F1 visual (can be different from actual compound)
	// 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet
	// F1 Classic – same as above
	// F2 – same as above

	VisualTyreCompound uint8
	TyresAgeLaps       uint8 // Age in laps of the current set of tyres

	RearLeftTyresDamage   uint8 // Tyre damage (percentage)
	RearRightTyresDamage  uint8
	FrontLeftTyresDamage  uint8
	FrontRightTyresDamage uint8

	FrontLeftWingDamage  uint8 // Front left wing damage (percentage)
	FrontRightWingDamage uint8 // Front right wing damage (percentage)
	RearWingDamage       uint8 // Rear wing damage (percentage)

	// Added Beta 3:
	DrsFault uint8 // Indicator for DRS fault, 0 = OK, 1 = fault

	EngineDamage  uint8 // Engine damage (percentage)
	GearBoxDamage uint8 // Gear box damage (percentage)

	// -1 = invalid/unknown, 0 = none, 1 = green
	// 2 = blue, 3 = yellow, 4 = red
	VehicleFiaFlags int8
	ErsStoreEnergy  float64 // ERS energy store in Joules

	// ERS deployment mode, 0 = none, 1 = medium
	// 2 = overtake, 3 = hotlap
	ErsDeployMode           uint8
	ErsHarvestedThisLapMGUK float64 // ERS energy harvested this lap by MGU-K
	ErsHarvestedThisLapMGUH float64 // ERS energy harvested this lap by MGU-H
	ErsDeployedThisLap      float64 // ERS energy deployed this lap
}

// CarStatusData details car statuses for all the cars in the race.
// It includes values such as the damage readings on the car.
type CarStatusData struct {
	Header Header

	CarStatusData CarStatusDataDetail
}

func NewCarStatusData(header packet.PacketHeader, p *packet.PacketCarStatusData) *CarStatusData {
	pk := p.CarStatusData[header.PlayerCarIndex]
	return &CarStatusData{
		Header: NewHeader(header),
		CarStatusData: CarStatusDataDetail{
			TractionControl:         pk.TractionControl,
			AntiLockBrakes:          pk.AntiLockBrakes,
			FuelMix:                 pk.FuelMix,
			FrontBrakeBias:          pk.FrontBrakeBias,
			PitLimiterStatus:        pk.PitLimiterStatus,
			FuelInTank:              pk.FuelInTank,
			FuelCapacity:            pk.FuelCapacity,
			FuelRemainingLaps:       pk.FuelRemainingLaps,
			MaxRPM:                  pk.MaxRPM,
			IdleRPM:                 pk.IdleRPM,
			MaxGears:                pk.MaxGears,
			DrsAllowed:              pk.DrsAllowed,
			DrsActivationDistance:   pk.DrsActivationDistance,
			RearLeftTyresWear:       pk.TyresWear[0],
			RearRightTyresWear:      pk.TyresWear[1],
			FrontLeftTyresWear:      pk.TyresWear[2],
			FrontRightTyresWear:     pk.TyresWear[3],
			ActualTyreCompound:      pk.ActualTyreCompound,
			VisualTyreCompound:      pk.VisualTyreCompound,
			TyresAgeLaps:            pk.TyresAgeLaps,
			RearLeftTyresDamage:     pk.TyresDamage[0],
			RearRightTyresDamage:    pk.TyresDamage[1],
			FrontLeftTyresDamage:    pk.TyresDamage[2],
			FrontRightTyresDamage:   pk.TyresDamage[3],
			FrontLeftWingDamage:     pk.FrontLeftWingDamage,
			FrontRightWingDamage:    pk.FrontRightWingDamage,
			RearWingDamage:          pk.RearWingDamage,
			DrsFault:                pk.DrsFault,
			EngineDamage:            pk.EngineDamage,
			GearBoxDamage:           pk.GearBoxDamage,
			VehicleFiaFlags:         pk.VehicleFiaFlags,
			ErsStoreEnergy:          pk.ErsStoreEnergy,
			ErsDeployMode:           pk.ErsDeployMode,
			ErsHarvestedThisLapMGUK: pk.ErsHarvestedThisLapMGUK,
			ErsHarvestedThisLapMGUH: pk.ErsHarvestedThisLapMGUH,
			ErsDeployedThisLap:      pk.ErsDeployedThisLap,
		},
	}
}

func (p *CarStatusData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
