package models

import (
	"strconv"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

type Header struct {
	PacketFormat     uint16  // 2020
	GameMajorVersion uint8   // Game major version - "X.00"
	GameMinorVersion uint8   // Game minor version - "1.XX"
	PacketVersion    uint8   // Version of this packet type, all start from 1
	PacketID         uint8   // Identifier for the packet type, see below
	SessionUID       string  // Unique identifier for the session
	SessionTime      float64 // Session timestamp
	FrameIdentifier  uint32  // Identifier for the frame the data was retrieved on
	PlayerCarIndex   uint8   // Index of player's car in the array

	// ADDED IN BETA 2:
	SecondaryPlayerCarIndex uint8 // Index of secondary player's car in the array (splitscreen)
	// 255 if no second player
}

func NewHeader(p packet.PacketHeader) Header {
	return Header{
		PacketFormat:            p.PacketFormat,
		GameMajorVersion:        p.GameMajorVersion,
		GameMinorVersion:        p.GameMinorVersion,
		PacketVersion:           p.PacketVersion,
		PacketID:                p.PacketID,
		SessionUID:              strconv.FormatUint(p.SessionUID, 10),
		SessionTime:             p.SessionTime,
		FrameIdentifier:         p.FrameIdentifier,
		PlayerCarIndex:          p.PlayerCarIndex,
		SecondaryPlayerCarIndex: p.SecondaryPlayerCarIndex,
	}
}
