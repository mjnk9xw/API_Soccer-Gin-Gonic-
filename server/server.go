package server

import (
	"Soccer-oauth2/err"
	"Soccer-oauth2/interfacemodel"
	"Soccer-oauth2/models"
	"Soccer-oauth2/serverstore"
)

// GetAllTour lay tat ca tour
func GetAllTour() (tours []interfacemodel.TourInfo, err error) {
	tours, err = serverstore.Tours.GetAll()
	return
}

// GetTourByID lay tour theo id
func GetTourByID(id string) (tour interfacemodel.TourInfo, err error) {
	tour = nil
	err = errSC.Success
	idInt, err := CheckTour(id)
	if err == errSC.Success {
		tour, err = serverstore.Tours.GetByID(idInt)
		return
	}
	return
}

// AddTour them 1 tour
func AddTour(t *models.Tour) (err error) {
	err = CheckTourName(t.Name)
	if err == errSC.Success {
		serverstore.Tours.Set(t)
		serverstore.TourIDToListRoundID.Set(t.GenIDKey())
		serverstore.TourIDToListPlayerID.Set(t.GenIDKey())
	}
	return
}

// UpdateTour chinh sua 1 tour
func UpdateTour(t *models.Tour) error {
	serverstore.Tours.Set(t)
	return errSC.Success
}

// DeleteTour xoa 1 tour
func DeleteTour(id string) (err error) {
	idInt, err := CheckTour(id)
	if err == errSC.Success {
		err = serverstore.Tours.Remove(idInt)
		serverstore.TourIDToListRoundID.Remove(id)
		serverstore.TourIDToListPlayerID.Remove(id)
		return
	}
	return
}

// GetAllRound lay tat ca round trong 1 tour
func GetAllRound(id string) (rounds []interfacemodel.RoundInfo, err error) {
	_, err = CheckTour(id)
	if err == errSC.Success {
		rounds, err = ToursByID(id)
	}
	return
}

// GetRoundByID lay 1 round theo id trong 1 tour
func GetRoundByID(idTour string, idRound string) (round interfacemodel.RoundInfo, err error) {
	round = nil
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		idRoundInt, err := CheckRound(idRound)
		if err == errSC.Success {
			round, err = serverstore.Rounds.GetByID(idRoundInt)
		}
	}
	return
}

// AddRound them 1 round trong 1 tour
func AddRound(id string, r *models.Round) (err error) {
	_, err = CheckTour(id)
	if err == errSC.Success {
		serverstore.Rounds.Set(r)
		serverstore.RoundIDToListTableID.Set(r.GetRoundID())
		serverstore.TourIDToListRoundID.SetRound(id, r.GetRoundID())
	}
	return
}

// UpdateRound chinh sua 1 round trong 1 tour
func UpdateRound(id string, r *models.Round) (err error) {
	_, err = CheckTour(id)
	if err == errSC.Success {
		serverstore.Rounds.Set(r)
	}
	return
}

// DeleteRound xoa 1 round trong 1 tour
func DeleteRound(idTour string, idRound string) (err error) {
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		IDRoundInt, err := CheckRound(idRound)
		if err == errSC.Success {
			serverstore.Rounds.Remove(IDRoundInt)
			serverstore.RoundIDToListTableID.Remove(IDRoundInt)
			serverstore.TourIDToListRoundID.RemoveRound(idTour, IDRoundInt)
		}
	}
	return
}

// GetAllTable lay tat ca bang dau trong 1 round
func GetAllTable(idTour, idRound string) (tables []interfacemodel.TableInfo, err error) {
	tables = nil
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		IDRoundInt, err := CheckRound(idRound)
		if err == errSC.Success {
			tables, err = RoundByID(IDRoundInt)
		}
	}
	return
}

// GetTableByID lay 1 bang dau theo id trong round
func GetTableByID(idTour, idRound, idTable string) (table interfacemodel.TableInfo, err error) {
	_, err = CheckTour(idTour)
	table = nil
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			IDTableInt, err := CheckTable(idTable)
			if err == errSC.Success {
				table, err = serverstore.Tables.GetByID(IDTableInt)
				lstPlayers, _ := PlayersInTable(IDTableInt)
				table.SetPlayers(lstPlayers)
				lstMatchs, _ := MatchInTable(IDTableInt)
				table.SetMatchs(lstMatchs)
			}
		}
	}
	return
}

// AddTable them 1 bang dau vao round
func AddTable(idTour, idRound string, t *models.Table) (err error) {
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		IDRoundInt, err := CheckRound(idRound)
		if err == errSC.Success {
			serverstore.Tables.Set(t)
			serverstore.RoundIDToListTableID.SetTable(IDRoundInt, t.GetTableID())
			serverstore.TableIDToListMatchID.Set(t.GetTableID())
			serverstore.TableIDToListPlayerID.Set(t.GetTableID())
		}
	}
	return
}

// UpdateTable chinh sua 1 bang dau trong round
func UpdateTable(idTour, idRound string, t *models.Table) (err error) {
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			serverstore.Tables.Set(t)
		}
	}
	return
}

// DeleteTable xoa 1 bang dau trong round
func DeleteTable(idTour, idRound, idTable string) (err error) {
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		IDRoundInt, err := CheckRound(idRound)
		if err == errSC.Success {
			IDTableInt, err := CheckTable(idTable)
			if err == errSC.Success {
				serverstore.Tables.Remove(IDTableInt)
				serverstore.RoundIDToListTableID.RemoveTable(IDRoundInt, IDTableInt)
				serverstore.TableIDToListMatchID.Remove(IDTableInt)
				serverstore.TableIDToListPlayerID.Remove(IDTableInt)
			}
		}
	}
	return
}

// GetAllPlayerInTable  lay tat ca doi bong trong 1 bang
func GetAllPlayerInTable(idTour, idRound, idTable string) (data []interfacemodel.PlayerInfo, err error) {
	data = nil
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			IDTableInt, err := CheckTable(idTable)
			if err == errSC.Success {
				data, err = PlayersInTable(IDTableInt)
			}
		}
	}
	return
}

// AddPlayer them 1 doi bong trong 1 bang
func AddPlayer(idTour, idRound, idTable string, p *models.Player) (err error) {
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			IDTableInt, err := CheckTable(idTable)
			if err == errSC.Success {
				err = CheckPlayerName(p.PlayerName)
				if err == errSC.Success {
					serverstore.Players.Set(p)
					serverstore.TourIDToListPlayerID.SetPlayer(idTour, p.GetPlayerID())
					serverstore.TableIDToListPlayerID.SetPlayer(IDTableInt, p.GetPlayerID())
				}
			}
		}
	}
	return
}

// UpdatePlayer chinh sua 1 doi bong trong 1 bang
func UpdatePlayer(idTour, idRound, idTable string, p *models.Player) (err error) {
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			_, err := CheckTable(idTable)
			if err == errSC.Success {
				serverstore.Players.Set(p)
			}
		}
	}
	return
}

// DeletePlayer xoa 1 doi bong trong 1 bang
func DeletePlayer(idTour, idRound, idTable, idPlayer string) (err error) {
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			IDTableInt, err := CheckTable(idTable)
			if err == errSC.Success {
				IDPlayerInt, err := CheckPlayer(idPlayer)
				if err == errSC.Success {
					serverstore.Players.Remove(IDPlayerInt)
					serverstore.TourIDToListPlayerID.RemovePlayer(idTour, IDPlayerInt)
					serverstore.TableIDToListPlayerID.RemovePlayer(IDTableInt, IDPlayerInt)
				}
			}
		}
	}
	return
}

// GetAllMatchInTable lay tat ca tran dau trong 1 bang
func GetAllMatchInTable(idTour, idRound, idTable string) (data []interfacemodel.MatchInfo, err error) {
	data = nil
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			IDTableInt, err := CheckTable(idTable)
			if err == errSC.Success {
				data, err = MatchInTable(IDTableInt)
			}
		}
	}
	return
}

// GetMatchByID lay 1 tran dau theo id trong 1 bang
func GetMatchByID(idTour, idRound, idTable, idMatch string) (match interfacemodel.MatchInfo, err error) {
	match = nil
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			_, err := CheckTable(idTable)
			if err == errSC.Success {
				IDMatchInt, err := CheckMatch(idMatch)
				if err == errSC.Success {
					match, err = serverstore.Matchs.GetByID(IDMatchInt)
				}
			}
		}
	}
	return
}

// AddMatch them 1 tran dau trong 1 bang
func AddMatch(idTour, idRound, idTable string, m *models.Match) (err error) {
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			IDTableInt, err := CheckTable(idTable)
			if err == errSC.Success {
				serverstore.Matchs.Set(m)
				serverstore.TableIDToListMatchID.SetMatch(IDTableInt, m.GetMatchID())
			}
		}
	}
	return
}

// UpdateMatch chinh sau 1 tran dau trong 1 bang
func UpdateMatch(idTour, idRound, idTable string, m *models.Match) (err error) {
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			_, err := CheckTable(idTable)
			if err == errSC.Success {
				serverstore.Matchs.Set(m)
			}
		}
	}
	return
}

// DeleteMatch xoa 1 tran dau trong 1 bang
func DeleteMatch(idTour, idRound, idTable, idMatch string) (err error) {
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		_, err := CheckRound(idRound)
		if err == errSC.Success {
			IDTableInt, err := CheckTable(idTable)
			if err == errSC.Success {
				IDMatchInt, err := CheckMatch(idMatch)
				if err == errSC.Success {
					serverstore.Matchs.Remove(IDMatchInt)
					serverstore.TableIDToListMatchID.RemoveMatch(IDTableInt, IDMatchInt)
				}
			}
		}
	}
	return
}

// GetPlayerAll lay tat ca doi bong trong 1 tour
func GetPlayerAll(idTour string) (players []interfacemodel.PlayerInfo, err error) {
	players = nil
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		players, err = PlayersInTour(idTour)
	}
	return
}

// GetPlayerInTour lay 1 doi bong theo id trong 1 tour
func GetPlayerInTour(idTour, idPlayer string) (player interfacemodel.PlayerInfo, err error) {
	player = nil
	_, err = CheckTour(idTour)
	if err == errSC.Success {
		IDPlayerInt, err := CheckPlayer(idPlayer)
		if err == errSC.Success {
			player, err = serverstore.Players.GetByID(IDPlayerInt)
		}
	}
	return
}
