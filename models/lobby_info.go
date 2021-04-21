package models

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/Tommy-42/f1-2020-go-telemetry/models/packet"
)

// LobbyInfoData details the players currently in a multiplayer lobby. It details each player’s selected car, any AI involved in the game and also the ready status of each of the participants.
type LobbyInfoData struct {
	Header    Header
	Timestamp time.Time

	//  specific data
	NumPlayers   uint8 // Number of players in the lobby data
	LobbyPlayers packet.LobbyInfoData
}

func NewLobbyInfoData(p *packet.PacketLobbyInfoData) *LobbyInfoData {
	return &LobbyInfoData{
		Header:       NewHeader(p.Header),
		Timestamp:    time.Now().UTC(),
		NumPlayers:   p.NumPlayers,
		LobbyPlayers: p.LobbyPlayers[p.Header.PlayerCarIndex],
	}
}

func (p *LobbyInfoData) ToJson() (*bytes.Reader, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
