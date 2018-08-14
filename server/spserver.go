package server

import (
	"Soccer-oauth2/err"
	"Soccer-oauth2/interfacemodel"
	"Soccer-oauth2/serverstore"
	"strconv"
)

// CheckIDInt kiem tra xem data co phai int khong
func CheckIDInt(id string) (idint int64, err error) {
	idint, err = strconv.ParseInt(id, 0, 64)
	if err == nil {
		err = errSC.Success
		return
	}
	err = errSC.ErrIDInt
	return
}

// CheckTourName kiem tra giai dau da co chua ko ?
func CheckTourName(name string) (err error) {
	names, _ := serverstore.Tours.GetName()
	err = errSC.Success
	for _, v := range names {
		if v == name {
			err = errSC.ErrTourName
			break
		}
	}
	return
}

// CheckPlayerName kiem tra dou bong da co chua ?
func CheckPlayerName(name string) (err error) {
	names, _ := serverstore.Players.GetName()
	err = errSC.Success
	for _, v := range names {
		if v == name {
			err = errSC.ErrPlayerName
			break
		}
	}
	return
}

// CheckTour kiem tra xem giai dau co ton tai chua ?
func CheckTour(id string) (idInt int64, err error) {
	idInt, err = CheckIDInt(id)
	if err == errSC.Success {
		_, err = serverstore.Tours.GetByID(idInt)
	}
	return
}

// CheckRound kiem tra xem vong dau co ton tai chua ?
func CheckRound(id string) (idInt int64, err error) {
	idInt, err = CheckIDInt(id)
	if err == errSC.Success {
		_, err = serverstore.Rounds.GetByID(idInt)
	}
	return
}

// CheckTable kiem tra xem bang dau co ton tai chua ?
func CheckTable(id string) (idInt int64, err error) {
	idInt, err = CheckIDInt(id)
	if err == errSC.Success {
		_, err = serverstore.Tables.GetByID(idInt)
	}
	return
}

// CheckPlayer kiem tra xem vong dau co ton tai chua ?
func CheckPlayer(id string) (idInt int64, err error) {
	idInt, err = CheckIDInt(id)
	if err == errSC.Success {
		_, err = serverstore.Players.GetByID(idInt)
	}
	return
}

// CheckMatch kiem tra xem vong dau co ton tai chua ?
func CheckMatch(id string) (idInt int64, err error) {
	idInt, err = CheckIDInt(id)
	if err == errSC.Success {
		_, err = serverstore.Matchs.GetByID(idInt)
	}
	return
}

// ToursByID lay tat ca round trong 1 tour
func ToursByID(id string) (rounds []interfacemodel.RoundInfo, err error) {
	roundsID, err := serverstore.TourIDToListRoundID.GetByID(id)
	rounds = make([]interfacemodel.RoundInfo, 0)
	for _, v := range roundsID {
		round, _ := serverstore.Rounds.GetByID(v)
		rounds = append(rounds, round)
	}
	return
}

// RoundByID lay tat ca table trong 1 round
func RoundByID(id int64) (tables []interfacemodel.TableInfo, err error) {
	tablesID, err := serverstore.RoundIDToListTableID.GetByID(id)
	tables = make([]interfacemodel.TableInfo, 0)
	for _, v := range tablesID {
		table, _ := serverstore.Tables.GetByID(v)
		tables = append(tables, table)
	}
	return
}

// PlayersInTour lay tat ca doi bong trong 1 tour
func PlayersInTour(id string) (players []interfacemodel.PlayerInfo, err error) {
	playersID, err := serverstore.TourIDToListPlayerID.GetByID(id)
	players = make([]interfacemodel.PlayerInfo, 0)
	for _, v := range playersID {
		player, _ := serverstore.Players.GetByID(v)
		players = append(players, player)
	}
	return
}

// PlayersInTable lay tat ca doi bong trong 1 table
func PlayersInTable(id int64) (players []interfacemodel.PlayerInfo, err error) {
	playersID, err := serverstore.TableIDToListPlayerID.GetByID(id)
	players = make([]interfacemodel.PlayerInfo, 0)
	for _, v := range playersID {
		player, _ := serverstore.Players.GetByID(v)
		players = append(players, player)
	}
	return
}

// MatchInTable lay tat ca tran dau trong 1 table
func MatchInTable(id int64) (matchs []interfacemodel.MatchInfo, err error) {
	matchsID, err := serverstore.TableIDToListMatchID.GetByID(id)
	matchs = make([]interfacemodel.MatchInfo, 0)
	for _, v := range matchsID {
		match, _ := serverstore.Matchs.GetByID(v)
		matchs = append(matchs, match)
	}
	return
}
