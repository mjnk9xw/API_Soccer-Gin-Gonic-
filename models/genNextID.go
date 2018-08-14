package models

var (
	CounterTourID   = int64(0)
	CounterRoundID  = int64(0)
	CounterTableID  = int64(0)
	CounterPlayerID = int64(0)
	CounterMatchID  = int64(0)
)

func genNextTourID() int64 {
	CounterTourID++
	return CounterTourID
}
func genNextRoundID() int64 {
	CounterRoundID++
	return CounterRoundID
}
func genNextTableID() int64 {
	CounterTableID++
	return CounterTableID
}
func genNextPlayerID() int64 {
	CounterPlayerID++
	return CounterPlayerID
}
func genNextMatchID() int64 {
	CounterMatchID++
	return CounterMatchID
}
