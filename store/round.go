package store

import (
	"Soccer-oauth2/err"
	"Soccer-oauth2/interfacemodel"
	"sync"
)

// RoundStore data Round
type RoundStore struct {
	sync.RWMutex
	data map[int64]interfacemodel.RoundInfo
}

// NewRoundStore Khoi tao data round
func NewRoundStore() *RoundStore {
	return &RoundStore{
		data: make(map[int64]interfacemodel.RoundInfo),
	}
}

// GetByID lay 1 round trong data
func (ts *RoundStore) GetByID(id int64) (tf interfacemodel.RoundInfo, err error) {
	ts.RLock()
	defer ts.RUnlock()
	if t, ok := ts.data[id]; ok {
		tf = t
		err = errSC.Success
		return
	}
	err = errSC.NotFound
	return

}

// Set them 1 round vao data
func (ts *RoundStore) Set(tf interfacemodel.RoundInfo) {
	ts.RLock()
	defer ts.RUnlock()
	ts.data[tf.GetRoundID()] = tf
	return
}

// Remove xoa 1 round trong data
func (ts *RoundStore) Remove(id int64) {
	ts.RLock()
	defer ts.RUnlock()
	delete(ts.data, id)
	return
}

// RoundIDtoTableIDStore RoundID to list TableID
type RoundIDtoTableIDStore struct {
	sync.RWMutex
	data map[int64][]int64
}

// NewRoundIDtoTableIDStore khoi tao RoundID to list TableID
func NewRoundIDtoTableIDStore() *RoundIDtoTableIDStore {
	return &RoundIDtoTableIDStore{
		data: make(map[int64][]int64),
	}
}

// GetByID lay 1 lst tableid theo round id
func (trID *RoundIDtoTableIDStore) GetByID(id int64) (lst []int64, err error) {
	trID.RLock()
	defer trID.RUnlock()
	err = errSC.Success
	if t, ok := trID.data[id]; ok {
		lst = t
		return
	}
	err = errSC.ErrIDRound
	return

}

// Set them 1 round id trong data nay
func (trID *RoundIDtoTableIDStore) Set(id int64) {
	trID.RLock()
	defer trID.RUnlock()
	trID.data[id] = nil
	return
}

// Remove xoa 1 round id trong data nay
func (trID *RoundIDtoTableIDStore) Remove(id int64) {
	trID.RLock()
	defer trID.RUnlock()
	delete(trID.data, id)
	return
}

// SetTable them 1 table id trong data nay
func (trID *RoundIDtoTableIDStore) SetTable(idRound int64, idTable int64) {
	trID.RLock()
	defer trID.RUnlock()
	trID.data[idRound] = append(trID.data[idRound], idTable)
	return
}

// RemoveTable xoa 1 table id trong data nay
func (trID *RoundIDtoTableIDStore) RemoveTable(idRound, idTable int64) error {
	trID.RLock()
	defer trID.RUnlock()
	for i, v := range trID.data[idRound] {
		if v == idTable {
			trID.data[idRound] = append(trID.data[idRound][:i], trID.data[idRound][i+1:]...)
			break
		}
	}
	return errSC.Success
}
