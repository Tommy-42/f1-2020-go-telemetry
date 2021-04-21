package models

import (
	"bytes"
	"encoding/json"
	"strconv"

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
	RearLeftSuspensionPosition       string // Note: All wheel arrays have the following order:
	RearRightSuspensionPosition      string // Note: All wheel arrays have the following order:
	FrontLeftSuspensionPosition      string // Note: All wheel arrays have the following order:
	FrontRightSuspensionPosition     string // Note: All wheel arrays have the following order:
	RearLeftSuspensionVelocity       string // RL, RR, FL, FR
	RearRightSuspensionVelocity      string // RL, RR, FL, FR
	FrontLeftSuspensionVelocity      string // RL, RR, FL, FR
	FrontRightSuspensionVelocity     string // RL, RR, FL, FR
	RearLeftSuspensionAcceleration   string // RL, RR, FL, FR
	RearRightSuspensionAcceleration  string // RL, RR, FL, FR
	FrontLeftSuspensionAcceleration  string // RL, RR, FL, FR
	FrontRightSuspensionAcceleration string // RL, RR, FL, FR
	RearLeftWheelSpeed               string // Speed of each wheel
	RearRightWheelSpeed              string // Speed of each wheel
	FrontLeftWheelSpeed              string // Speed of each wheel
	FrontRightWheelSpeed             string // Speed of each wheel
	RearLeftWheelSlip                string // Slip ratio for each wheel
	RearRightWheelSlip               string // Slip ratio for each wheel
	FrontLeftWheelSlip               string // Slip ratio for each wheel
	FrontRightWheelSlip              string // Slip ratio for each wheel
	LocalVelocityX                   string // Velocity in local space
	LocalVelocityY                   string // Velocity in local space
	LocalVelocityZ                   string // Velocity in local space
	AngularVelocityX                 string // Angular velocity x-component
	AngularVelocityY                 string // Angular velocity y-component
	AngularVelocityZ                 string // Angular velocity z-component
	AngularAccelerationX             string // Angular velocity x-component
	AngularAccelerationY             string // Angular velocity y-component
	AngularAccelerationZ             string // Angular velocity z-component
	FrontWheelsAngle                 string // Current front wheels angle in radians
}

func NewMotionData(p *packet.PacketMotionData) *MotionData {
	return &MotionData{
		Header:                           NewHeader(p.Header),
		CarMotionData:                    p.CarMotionData[p.Header.PlayerCarIndex],
		RearLeftSuspensionPosition:       strconv.FormatFloat(float64(p.SuspensionPosition[0]), 'f', 4, 64),
		RearRightSuspensionPosition:      strconv.FormatFloat(float64(p.SuspensionPosition[1]), 'f', 4, 64),
		FrontLeftSuspensionPosition:      strconv.FormatFloat(float64(p.SuspensionPosition[2]), 'f', 4, 64),
		FrontRightSuspensionPosition:     strconv.FormatFloat(float64(p.SuspensionPosition[3]), 'f', 4, 64),
		RearLeftSuspensionVelocity:       strconv.FormatFloat(float64(p.SuspensionVelocity[0]), 'f', 4, 64),
		RearRightSuspensionVelocity:      strconv.FormatFloat(float64(p.SuspensionVelocity[1]), 'f', 4, 64),
		FrontLeftSuspensionVelocity:      strconv.FormatFloat(float64(p.SuspensionVelocity[2]), 'f', 4, 64),
		FrontRightSuspensionVelocity:     strconv.FormatFloat(float64(p.SuspensionVelocity[3]), 'f', 4, 64),
		RearLeftSuspensionAcceleration:   strconv.FormatFloat(float64(p.SuspensionAcceleration[0]), 'f', 4, 64),
		RearRightSuspensionAcceleration:  strconv.FormatFloat(float64(p.SuspensionAcceleration[1]), 'f', 4, 64),
		FrontLeftSuspensionAcceleration:  strconv.FormatFloat(float64(p.SuspensionAcceleration[2]), 'f', 4, 64),
		FrontRightSuspensionAcceleration: strconv.FormatFloat(float64(p.SuspensionAcceleration[3]), 'f', 4, 64),
		RearLeftWheelSpeed:               strconv.FormatFloat(float64(p.WheelSpeed[0]), 'f', 4, 64),
		RearRightWheelSpeed:              strconv.FormatFloat(float64(p.WheelSpeed[1]), 'f', 4, 64),
		FrontLeftWheelSpeed:              strconv.FormatFloat(float64(p.WheelSpeed[2]), 'f', 4, 64),
		FrontRightWheelSpeed:             strconv.FormatFloat(float64(p.WheelSpeed[3]), 'f', 4, 64),
		RearLeftWheelSlip:                strconv.FormatFloat(float64(p.WheelSlip[0]), 'f', 4, 64),
		RearRightWheelSlip:               strconv.FormatFloat(float64(p.WheelSlip[1]), 'f', 4, 64),
		FrontLeftWheelSlip:               strconv.FormatFloat(float64(p.WheelSlip[2]), 'f', 4, 64),
		FrontRightWheelSlip:              strconv.FormatFloat(float64(p.WheelSlip[3]), 'f', 4, 64),
		LocalVelocityX:                   strconv.FormatFloat(float64(p.LocalVelocityX), 'f', 4, 64),
		LocalVelocityY:                   strconv.FormatFloat(float64(p.LocalVelocityY), 'f', 4, 64),
		LocalVelocityZ:                   strconv.FormatFloat(float64(p.LocalVelocityZ), 'f', 4, 64),
		AngularVelocityX:                 strconv.FormatFloat(float64(p.AngularVelocityX), 'f', 4, 64),
		AngularVelocityY:                 strconv.FormatFloat(float64(p.AngularVelocityY), 'f', 4, 64),
		AngularVelocityZ:                 strconv.FormatFloat(float64(p.AngularVelocityZ), 'f', 4, 64),
		AngularAccelerationX:             strconv.FormatFloat(float64(p.AngularAccelerationX), 'f', 4, 64),
		AngularAccelerationY:             strconv.FormatFloat(float64(p.AngularAccelerationY), 'f', 4, 64),
		AngularAccelerationZ:             strconv.FormatFloat(float64(p.AngularAccelerationZ), 'f', 4, 64),
		FrontWheelsAngle:                 strconv.FormatFloat(float64(p.FrontWheelsAngle), 'f', 4, 64),
	}
}

func (p *MotionData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
