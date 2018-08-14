package models

type RoundTime struct {
	StartTime string `json:"start"`
	EndTime   string `json:"end"`
}

type Round struct {
	RoundID   int64      `json:"id"`
	RoundName string     `json:"name"`
	TimeR     *RoundTime `json:"time"`
}

func (t *Round) SetRoundID() {
	t.RoundID = genNextRoundID()
}
func (t *Round) GetRoundID() int64 {
	return t.RoundID
}
