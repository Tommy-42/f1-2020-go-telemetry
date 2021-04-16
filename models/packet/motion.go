package packet

// CarMotionData data for the car being driven
type CarMotionData struct {
	WorldPositionX     float64 // World space X position
	WorldPositionY     float64 // World space Y position
	WorldPositionZ     float64 // World space Z position
	WorldVelocityX     float64 // Velocity in world space X
	WorldVelocityY     float64 // Velocity in world space Y
	WorldVelocityZ     float64 // Velocity in world space Z
	WorldForwardDirX   int16   // World space forward X direction (normalised)
	WorldForwardDirY   int16   // World space forward Y direction (normalised)
	WorldForwardDirZ   int16   // World space forward Z direction (normalised)
	WorldRightDirX     int16   // World space right X direction (normalised)
	WorldRightDirY     int16   // World space right Y direction (normalised)
	WorldRightDirZ     int16   // World space right Z direction (normalised)
	GForceLateral      float64 // Lateral G-Force component
	GForceLongitudinal float64 // Longitudinal G-Force component
	GForceVertical     float64 // Vertical G-Force component
	Yaw                float64 // Yaw angle in radians
	Pitch              float64 // Pitch angle in radians
	Roll               float64 // Roll angle in radians
}

// PacketMotionData the motion packet gives physics data for all the cars being driven.
// There is additional data for the car being driven with the goal of being able to drive a motion platform setup.
//
// N.B. For the normalised vectors below,
// to convert to float values divide by 32767.0f – 16-bit signed values are used to pack the data and
// on the assumption that direction values are always between -1.0f and 1.0f.
//
// Frequency: Rate as specified in menus
// Size: 1464 bytes (Packet size updated in Beta 3)
// Version: 1
type PacketMotionData struct {
	CarMotionData [22]CarMotionData // Data for all cars on track

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
