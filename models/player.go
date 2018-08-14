package models

type Player struct {
	PlayerID      int64  `json:"id"`
	PlayerName    string `json:"name"`
	ComeNextRound bool   `json:"next"`
}

func (p *Player) SetPlayerID() {
	p.PlayerID = genNextPlayerID()
}
func (p *Player) GetPlayerID() int64 {
	return p.PlayerID
}
func (p *Player) GetPlayerName() string {
	return p.PlayerName
}
