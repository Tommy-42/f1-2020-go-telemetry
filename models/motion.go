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
	RearLeftSuspensionPosition       float64 // Note: All wheel arrays have the following order:
	RearRightSuspensionPosition      float64 // Note: All wheel arrays have the following order:
	FrontLeftSuspensionPosition      float64 // Note: All wheel arrays have the following order:
	FrontRightSuspensionPosition     float64 // Note: All wheel arrays have the following order:
	RearLeftSuspensionVelocity       float64 // RL, RR, FL, FR
	RearRightSuspensionVelocity      float64 // RL, RR, FL, FR
	FrontLeftSuspensionVelocity      float64 // RL, RR, FL, FR
	FrontRightSuspensionVelocity     float64 // RL, RR, FL, FR
	RearLeftSuspensionAcceleration   float64 // RL, RR, FL, FR
	RearRightSuspensionAcceleration  float64 // RL, RR, FL, FR
	FrontLeftSuspensionAcceleration  float64 // RL, RR, FL, FR
	FrontRightSuspensionAcceleration float64 // RL, RR, FL, FR
	RearLeftWheelSpeed               float64 // Speed of each wheel
	RearRightWheelSpeed              float64 // Speed of each wheel
	FrontLeftWheelSpeed              float64 // Speed of each wheel
	FrontRightWheelSpeed             float64 // Speed of each wheel
	RearLeftWheelSlip                float64 // Slip ratio for each wheel
	RearRightWheelSlip               float64 // Slip ratio for each wheel
	FrontLeftWheelSlip               float64 // Slip ratio for each wheel
	FrontRightWheelSlip              float64 // Slip ratio for each wheel
	LocalVelocityX                   float64 // Velocity in local space
	LocalVelocityY                   float64 // Velocity in local space
	LocalVelocityZ                   float64 // Velocity in local space
	AngularVelocityX                 float64 // Angular velocity x-component
	AngularVelocityY                 float64 // Angular velocity y-component
	AngularVelocityZ                 float64 // Angular velocity z-component
	AngularAccelerationX             float64 // Angular velocity x-component
	AngularAccelerationY             float64 // Angular velocity y-component
	AngularAccelerationZ             float64 // Angular velocity z-component
	FrontWheelsAngle                 float64 // Current front wheels angle in radians
}

func NewMotionData(header packet.PacketHeader, p *packet.PacketMotionData) *MotionData {
	return &MotionData{
		Header:                           NewHeader(header),
		CarMotionData:                    p.CarMotionData[header.PlayerCarIndex],
		RearLeftSuspensionPosition:       p.SuspensionPosition[0],
		RearRightSuspensionPosition:      p.SuspensionPosition[1],
		FrontLeftSuspensionPosition:      p.SuspensionPosition[2],
		FrontRightSuspensionPosition:     p.SuspensionPosition[3],
		RearLeftSuspensionVelocity:       p.SuspensionVelocity[0],
		RearRightSuspensionVelocity:      p.SuspensionVelocity[1],
		FrontLeftSuspensionVelocity:      p.SuspensionVelocity[2],
		FrontRightSuspensionVelocity:     p.SuspensionVelocity[3],
		RearLeftSuspensionAcceleration:   p.SuspensionAcceleration[0],
		RearRightSuspensionAcceleration:  p.SuspensionAcceleration[1],
		FrontLeftSuspensionAcceleration:  p.SuspensionAcceleration[2],
		FrontRightSuspensionAcceleration: p.SuspensionAcceleration[3],
		RearLeftWheelSpeed:               p.WheelSpeed[0],
		RearRightWheelSpeed:              p.WheelSpeed[1],
		FrontLeftWheelSpeed:              p.WheelSpeed[2],
		FrontRightWheelSpeed:             p.WheelSpeed[3],
		RearLeftWheelSlip:                p.WheelSlip[0],
		RearRightWheelSlip:               p.WheelSlip[1],
		FrontLeftWheelSlip:               p.WheelSlip[2],
		FrontRightWheelSlip:              p.WheelSlip[3],
		LocalVelocityX:                   p.LocalVelocityX,
		LocalVelocityY:                   p.LocalVelocityY,
		LocalVelocityZ:                   p.LocalVelocityZ,
		AngularVelocityX:                 p.AngularVelocityX,
		AngularVelocityY:                 p.AngularVelocityY,
		AngularVelocityZ:                 p.AngularVelocityZ,
		AngularAccelerationX:             p.AngularAccelerationX,
		AngularAccelerationY:             p.AngularAccelerationY,
		AngularAccelerationZ:             p.AngularAccelerationZ,
		FrontWheelsAngle:                 p.FrontWheelsAngle,
	}
}

func (p *MotionData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
