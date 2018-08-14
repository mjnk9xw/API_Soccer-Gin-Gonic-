package interfacemodel

type (
	TourInfo interface {
		GenIDKey() string
		GetTourID() int64
		SetTourID()
		GetTourName() string
	}

	RoundInfo interface {
		SetRoundID()
		GetRoundID() int64
	}

	TableInfo interface {
		SetTableID()
		GetTableID() int64
		SetPlayers(p []PlayerInfo)
		SetMatchs(m []MatchInfo)
	}

	PlayerInfo interface {
		SetPlayerID()
		GetPlayerID() int64
		GetPlayerName() string
	}

	MatchInfo interface {
		SetMatchID()
		GetMatchID() int64
	}
)
