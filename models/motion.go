package models

import (
	"bytes"
	"encoding/json"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// MotionData the motion packet gives physics data for all the cars being driven.
// There is additional data for the car being driven with the goal of being able to drive a motion platform setup.
//
// N.B. For the normalised vectors below,
// to convert to float values divide by 32767.0f – 16-bit signed values are used to pack the data and
// on the assumption that direction values are always between -1.0f and 1.0f.
type MotionData struct {
	Header Header

	CarMotionData packet.CarMotionData // Data for all cars on track

	// Extra player car ONLY data
	SuspensionPosition     [4]float64 // Note: All wheel arrays have the following order:
	SuspensionVelocity     [4]float64 // RL, RR, FL, FR
	SuspensionAcceleration [4]float64 // RL, RR, FL, FR
	WheelSpeed             [4]float64 // Speed of each wheel
	WheelSlip              [4]float64 // Slip ratio for each wheel
	LocalVelocityX         float64    // Velocity in local space
	LocalVelocityY         float64    // Velocity in local space
	LocalVelocityZ         float64    // Velocity in local space
	AngularVelocityX       float64    // Angular velocity x-component
	AngularVelocityY       float64    // Angular velocity y-component
	AngularVelocityZ       float64    // Angular velocity z-component
	AngularAccelerationX   float64    // Angular velocity x-component
	AngularAccelerationY   float64    // Angular velocity y-component
	AngularAccelerationZ   float64    // Angular velocity z-component
	FrontWheelsAngle       float64    // Current front wheels angle in radians
}

func NewMotionData(p *packet.PacketMotionData) *MotionData {
	if p.Header.PlayerCarIndex > 21 {
		return nil
	}

	return &MotionData{
		Header:                 NewHeader(p.Header),
		CarMotionData:          p.CarMotionData[p.Header.PlayerCarIndex],
		SuspensionPosition:     p.SuspensionPosition,
		SuspensionVelocity:     p.SuspensionVelocity,
		SuspensionAcceleration: p.SuspensionAcceleration,
		WheelSpeed:             p.WheelSpeed,
		WheelSlip:              p.WheelSlip,
		LocalVelocityX:         p.LocalVelocityX,
		LocalVelocityY:         p.LocalVelocityY,
		LocalVelocityZ:         p.LocalVelocityZ,
		AngularVelocityX:       p.AngularVelocityX,
		AngularVelocityY:       p.AngularVelocityY,
		AngularVelocityZ:       p.AngularVelocityZ,
		AngularAccelerationX:   p.AngularAccelerationX,
		AngularAccelerationY:   p.AngularAccelerationY,
		AngularAccelerationZ:   p.AngularAccelerationZ,
		FrontWheelsAngle:       p.FrontWheelsAngle,
	}
}

func (p *MotionData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
