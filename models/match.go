package models

type MatchTime struct {
	StartTime string `json:"start"`
}

type Match struct {
	MatchID        int64      `json:"id"`
	PlayerOneID    int64      `json:"PlayerOneID"`
	PlayerTwoID    int64      `json:"PlayerTwoID"`
	PlayerOneGoals int        `json:"PlayerOneGoals"`
	PlayerTwoGoals int        `json:"PlayerTwoGoals"`
	Time           *MatchTime `json:"time"`
}

func (t *Match) SetMatchID() {
	t.MatchID = genNextMatchID()
}
func (t *Match) GetMatchID() int64 {
	return t.MatchID
}
