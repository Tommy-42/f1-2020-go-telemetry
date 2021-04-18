package models

import (
	"bytes"
	"encoding/json"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// PacketEventData gives details of events that happen during the course of a session.
type EventData struct {
	Header Header

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
	EventStringCode string

	// EventDetails - should be interpreted differently for each type
	EventDetails packet.EventDataDetails
}

func NewEventData(p *packet.PacketEventData) *EventData {
	return &EventData{
		Header:          NewHeader(p.Header),
		EventStringCode: p.EventStringCode,
		EventDetails:    p.EventDetails,
	}
}

func (p *EventData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
