package udppacket

// CarSetupData ...
type CarSetupData struct {
	FrontWing              uint8   // Front wing aero
	RearWing               uint8   // Rear wing aero
	OnThrottle             uint8   // Differential adjustment on throttle (percentage)
	OffThrottle            uint8   // Differential adjustment off throttle (percentage)
	FrontCamber            float64 // Front camber angle (suspension geometry)
	RearCamber             float64 // Rear camber angle (suspension geometry)
	FrontToe               float64 // Front toe angle (suspension geometry)
	RearToe                float64 // Rear toe angle (suspension geometry)
	FrontSuspension        uint8   // Front suspension
	RearSuspension         uint8   // Rear suspension
	FrontAntiRollBar       uint8   // Front anti-roll bar
	RearAntiRollBar        uint8   // Front anti-roll bar
	FrontSuspensionHeight  uint8   // Front ride height
	RearSuspensionHeight   uint8   // Rear ride height
	BrakePressure          uint8   // Brake pressure (percentage)
	BrakeBias              uint8   // Brake bias (percentage)
	RearLeftTyrePressure   float64 // Rear left tyre pressure (PSI)
	RearRightTyrePressure  float64 // Rear right tyre pressure (PSI)
	FrontLeftTyrePressure  float64 // Front left tyre pressure (PSI)
	FrontRightTyrePressure float64 // Front right tyre pressure (PSI)
	Ballast                uint8   // Ballast
	FuelLoad               float64 // Fuel load
}

// PacketCarSetupData details the car setups for each vehicle in the session.
// Note that in multiplayer games, other player cars will appear as blank, you will only be able to see your car setup and AI cars.
//
// Frequency: 2 per second
// Size: 1102 bytes (Packet size updated in Beta 3)
// Version: 1
type PacketCarSetupData struct {
	Header PacketHeader // Header

	CarSetups [22]CarSetupData
}
