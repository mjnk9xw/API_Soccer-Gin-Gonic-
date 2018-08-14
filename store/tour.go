package store

import (
	"Soccer-oauth2/err"
	"Soccer-oauth2/interfacemodel"
	"sync"
)

// TourStore data Tour
type TourStore struct {
	sync.RWMutex
	data map[int64]interfacemodel.TourInfo
}

// NewTourStore khoi tao data Tour
func NewTourStore() *TourStore {
	return &TourStore{
		data: make(map[int64]interfacemodel.TourInfo),
	}
}

// GetByID lay 1 TourInfo trong data
func (ts *TourStore) GetByID(id int64) (tf interfacemodel.TourInfo, err error) {
	ts.RLock()
	defer ts.RUnlock()
	if t, ok := ts.data[id]; ok {
		tf = t
		err = errSC.Success
		return
	}
	err = errSC.ErrIDTour
	return
}

// GetAll lay toan bo Tour
func (ts *TourStore) GetAll() (lst []interfacemodel.TourInfo, err error) {
	ts.RLock()
	defer ts.RUnlock()
	err = errSC.Success
	lst = make([]interfacemodel.TourInfo, 0)
	for _, v := range ts.data {
		lst = append(lst, v)
	}
	return
}

// GetName lay toan bo Name Tour
func (ts *TourStore) GetName() (lst []string, err error) {
	ts.RLock()
	defer ts.RUnlock()
	err = errSC.Success
	lst = make([]string, 0)
	for _, v := range ts.data {
		lst = append(lst, v.GetTourName())
	}
	return
}

// Set them 1 TourInfo
func (ts *TourStore) Set(tf interfacemodel.TourInfo) {
	ts.RLock()
	defer ts.RUnlock()
	ts.data[tf.GetTourID()] = tf
	return
}

// Remove xoa 1 TourInfo
func (ts *TourStore) Remove(id int64) error {
	ts.RLock()
	defer ts.RUnlock()
	_, err := ts.data[id]
	if err {
		delete(ts.data, id)
		return errSC.Success
	}
	return errSC.NotFound
}

// TourIDtoRoundIDStore TourID to list Round ID
type TourIDtoRoundIDStore struct {
	sync.RWMutex
	data map[string][]int64
}

// NewTourIDtoRoundIDStore khoi tao TourID to list Round ID
func NewTourIDtoRoundIDStore() *TourIDtoRoundIDStore {
	return &TourIDtoRoundIDStore{
		data: make(map[string][]int64),
	}
}

// GetByID lay 1 lst RoundID
func (trID *TourIDtoRoundIDStore) GetByID(id string) (lst []int64, err error) {
	trID.RLock()
	defer trID.RUnlock()
	err = errSC.Success
	if t, ok := trID.data[id]; ok {
		lst = t
		return
	}
	err = errSC.NotFound
	return

}

// Set them 1 TourID
func (trID *TourIDtoRoundIDStore) Set(id string) {
	trID.RLock()
	defer trID.RUnlock()
	trID.data[id] = nil
	return
}

// SetRound Them 1 RoundID
func (trID *TourIDtoRoundIDStore) SetRound(idTour string, idRound int64) {
	trID.RLock()
	defer trID.RUnlock()
	trID.data[idTour] = append(trID.data[idTour], idRound)
	return
}

// Remove Xoa 1 TourID
func (trID *TourIDtoRoundIDStore) Remove(id string) {
	trID.RLock()
	defer trID.RUnlock()
	delete(trID.data, id)
	return
}

// RemoveRound xoa 1 RoundID
func (trID *TourIDtoRoundIDStore) RemoveRound(idTour string, idRound int64) error {
	trID.RLock()
	defer trID.RUnlock()
	for i, v := range trID.data[idTour] {
		if v == idRound {
			trID.data[idTour] = append(trID.data[idTour][:i], trID.data[idTour][i+1:]...)
			break
		}
	}
	return errSC.Success
}

//TourIDtoPlayerIDStore TourID to list Player
type TourIDtoPlayerIDStore struct {
	sync.RWMutex
	data map[string][]int64
}

// NewTourIDtoPlayerIDStore khoi tao data TourID to list Player
func NewTourIDtoPlayerIDStore() *TourIDtoPlayerIDStore {
	return &TourIDtoPlayerIDStore{
		data: make(map[string][]int64),
	}
}

// GetByID lay 1 lst PlayerID
func (trID *TourIDtoPlayerIDStore) GetByID(id string) (lst []int64, err error) {
	trID.RLock()
	defer trID.RUnlock()

	if t, ok := trID.data[id]; ok {
		lst = t
		return
	}
	err = errSC.NotFound
	return

}

// Set Them 1 TourID
func (trID *TourIDtoPlayerIDStore) Set(id string) {
	trID.RLock()
	defer trID.RUnlock()
	trID.data[id] = nil
	return
}

// SetPlayer Them 1 PlayerID
func (trID *TourIDtoPlayerIDStore) SetPlayer(idTour string, idPlayer int64) {
	trID.RLock()
	defer trID.RUnlock()
	trID.data[idTour] = append(trID.data[idTour], idPlayer)
	return
}

// Remove Xoa 1 TourID (key trong map)
func (trID *TourIDtoPlayerIDStore) Remove(id string) {
	trID.RLock()
	defer trID.RUnlock()
	delete(trID.data, id)
	return
}

// RemovePlayer Xoa 1 PlayerID trong map (value)
func (trID *TourIDtoPlayerIDStore) RemovePlayer(idTour string, idPlayer int64) error {
	trID.RLock()
	defer trID.RUnlock()
	for i, v := range trID.data[idTour] {
		if v == idPlayer {
			trID.data[idTour] = append(trID.data[idTour][:i], trID.data[idTour][i+1:]...)
			break
		}
	}
	return errSC.Success
}
