package serverstore

import (
	"Soccer-oauth2/store"
)

var (
	Tours                 = store.NewTourStore()
	Rounds                = store.NewRoundStore()
	Tables                = store.NewTableStore()
	Players               = store.NewPlayerStore()
	Matchs                = store.NewMatchStore()
	TourIDToListRoundID   = store.NewTourIDtoRoundIDStore()
	TourIDToListPlayerID  = store.NewTourIDtoPlayerIDStore()
	RoundIDToListTableID  = store.NewRoundIDtoTableIDStore()
	TableIDToListMatchID  = store.NewTableIDtoMatchIDStore()
	TableIDToListPlayerID = store.NewTableIDtoPlayerIDStore()
)
