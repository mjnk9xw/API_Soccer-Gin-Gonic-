package models

import (
	"Soccer-oauth2/interfacemodel"
)

type Table struct {
	TableID   int64                       `json:"id"`
	TableName string                      `json:"name"`
	Players   []interfacemodel.PlayerInfo `json:"players"`
	Matchs    []interfacemodel.MatchInfo  `json:"matchs"`
}

func (t *Table) SetTableID() {
	t.TableID = genNextTableID()
}
func (t *Table) GetTableID() int64 {
	return t.TableID
}
func (t *Table) SetPlayers(p []interfacemodel.PlayerInfo) {
	t.Players = make([]interfacemodel.PlayerInfo, 0)
	t.Players = append(t.Players, p...)
}
func (t *Table) SetMatchs(m []interfacemodel.MatchInfo) {
	t.Matchs = make([]interfacemodel.MatchInfo, 0)
	t.Matchs = append(t.Matchs, m...)
}
