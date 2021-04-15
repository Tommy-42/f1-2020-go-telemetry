package packet

import (
	"encoding/json"
)

// LobbyInfoData ...
type LobbyInfoData struct {
	AIControlled uint8 // Whether the vehicle is AI (1) or Human (0) controlled
	TeamID       uint8 // Team id - see appendix (255 if no team currently selected)
	Nationality  uint8 // Nationality of the driver
	// Name of participant in UTF-8 format – null terminated
	// Will be truncated with ... (U+2026) if too long
	Name        string
	ReadyStatus uint8 // 0 = not ready, 1 = ready, 2 = spectating
}

// PacketLobbyInfoData details the players currently in a multiplayer lobby. It details each player’s selected car, any AI involved in the game and also the ready status of each of the participants.
//
// Frequency: Two every second when in the lobby
// Size: 1169 bytes (Packet size updated in Beta 3)
// Version: 1
type PacketLobbyInfoData struct {
	Header PacketHeader // Header

	// Packet specific data
	NumPlayers   uint8 // Number of players in the lobby data
	LobbyPlayers [22]LobbyInfoData
}

func (p *PacketLobbyInfoData) Read(receiver []byte) (n int, err error) {
	data, err := json.Marshal(p)
	copy(receiver, data)
	n = len(data)
	return
}
