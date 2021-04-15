package packet

import (
	"encoding/json"
)

// EventDataDetails the event details packet is different for each type of event.
// Make sure only the correct type is interpreted.
type EventDataDetails interface{}

// FastestLap ...
type FastestLap struct {
	VehicleIdx uint8   // Vehicle index of car achieving fastest lap
	LapTime    float64 // Lap time is in seconds
}

// Retirement ...
type Retirement struct {
	VehicleIdx uint8 // Vehicle index of car retiring
}

// TeamMateInPits ...
type TeamMateInPits struct {
	VehicleIdx uint8 // Vehicle index of team mate
}

// RaceWinner ...
type RaceWinner struct {
	VehicleIdx uint8 // Vehicle index of the race winner
}

// Penalty ...
type Penalty struct {
	PenaltyType      uint8 // Penalty type – see Appendices
	InfringementType uint8 // Infringement type – see Appendices
	VehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 // Vehicle index of the other car involved
	Time             uint8 // Time gained, or time spent doing action in seconds
	LapNum           uint8 // Lap the penalty occurred on
	PlacesGained     uint8 // Number of places gained by this
}

// SpeedTrap ...
type SpeedTrap struct {
	VehicleIdx uint8   // Vehicle index of the vehicle triggering speed trap
	Speed      float64 // Top speed achieved in kilometres per hour
}

// PacketEventData gives details of events that happen during the course of a session.
//
// Frequency: When the event occurs
// Size: 35 bytes (Packet size updated in Beta 3)
// Version: 1
type PacketEventData struct {
	Header PacketHeader // Header

	/*
		EventStringCodes

		Event 					| Code 		| Description
		Session Started 		| “SSTA” 	| Sent when the session starts
		Session Ended 			| “SEND” 	| Sent when the session ends
		Fastest Lap 			| “FTLP” 	| When a driver achieves the fastest lap
		Retirement 				| “RTMT” 	| When a driver retires
		DRS enabled 			| “DRSE” 	| Race control have enabled DRS
		DRS disabled 			| “DRSD” 	| Race control have disabled DRS
		Team mate in pits 		| “TMPT” 	| Your team mate has entered the pits
		Chequered flag 			| “CHQF” 	| The chequered flag has been waved
		Race Winner 			| “RCWN” 	| The race winner is announced
		Penalty Issued 			| “PENA” 	| A penalty has been issued – details in event
		Speed Trap Triggered 	| “SPTP” 	| Speed trap has been triggered by fastest speed
	*/
	EventStringCode [4]uint8

	// EventDetails - should be interpreted differently for each type
	EventDetails EventDataDetails
}

func (p *PacketEventData) Read(receiver []byte) (n int, err error) {
	data, err := json.Marshal(p)
	copy(receiver, data)
	n = len(data)
	return
}
