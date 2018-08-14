package store

import (
	"Soccer-oauth2/err"
	"Soccer-oauth2/interfacemodel"
	"sync"
)

// TableStore data table
type TableStore struct {
	sync.RWMutex
	data map[int64]interfacemodel.TableInfo
}

// NewTableStore khoi tao data
func NewTableStore() *TableStore {
	return &TableStore{
		data: make(map[int64]interfacemodel.TableInfo),
	}
}

// GetByID lay 1 table theo id trong data
func (ts *TableStore) GetByID(id int64) (tf interfacemodel.TableInfo, err error) {
	ts.RLock()
	defer ts.RUnlock()
	if t, ok := ts.data[id]; ok {
		tf = t
		err = errSC.Success
		return
	}
	err = errSC.ErrIDTable
	return

}

// Set them 1 table vao data
func (ts *TableStore) Set(tf interfacemodel.TableInfo) {
	ts.RLock()
	defer ts.RUnlock()
	ts.data[tf.GetTableID()] = tf
	return
}

// Remove xoa 1 table trong data
func (ts *TableStore) Remove(id int64) {
	ts.RLock()
	defer ts.RUnlock()
	delete(ts.data, id)
	return
}

// TableIDtoMatchIDStore Table ID to list MatchID
type TableIDtoMatchIDStore struct {
	sync.RWMutex
	data map[int64][]int64
}

// NewTableIDtoMatchIDStore khoi tao data Table ID to list MatchID
func NewTableIDtoMatchIDStore() *TableIDtoMatchIDStore {
	return &TableIDtoMatchIDStore{
		data: make(map[int64][]int64),
	}
}

// GetByID lay 1 lst id Match trong table
func (trID *TableIDtoMatchIDStore) GetByID(id int64) (lst []int64, err error) {
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

// Set them 1 table id
func (trID *TableIDtoMatchIDStore) Set(id int64) {
	trID.RLock()
	defer trID.RUnlock()
	trID.data[id] = nil
	return
}

// Remove xoa 1 table id
func (trID *TableIDtoMatchIDStore) Remove(id int64) {
	trID.RLock()
	defer trID.RUnlock()
	delete(trID.data, id)
	return
}

// SetMatch them 1 Match id vao data TableIDtoMatchIDStore
func (trID *TableIDtoMatchIDStore) SetMatch(idTable int64, idMatch int64) {
	trID.RLock()
	defer trID.RUnlock()
	trID.data[idTable] = append(trID.data[idTable], idMatch)
	return
}

// RemoveMatch xoa 1 Match ID trong data TableIDtoMatchIDStore
func (trID *TableIDtoMatchIDStore) RemoveMatch(idTable int64, idMatch int64) error {
	trID.RLock()
	defer trID.RUnlock()
	for i, v := range trID.data[idTable] {
		if v == idMatch {
			trID.data[idTable] = append(trID.data[idTable][:i], trID.data[idTable][i+1:]...)
			break
		}
	}
	return errSC.Success
}

// TableIDtoPlayerIDStore Table ID to list Player
type TableIDtoPlayerIDStore struct {
	sync.RWMutex
	data map[int64][]int64
}

// NewTableIDtoPlayerIDStore khoi tao data Table ID to list Player
func NewTableIDtoPlayerIDStore() *TableIDtoPlayerIDStore {
	return &TableIDtoPlayerIDStore{
		data: make(map[int64][]int64),
	}
}

// GetByID lay 1 idplayer trong  data Table ID to list Player
func (trID *TableIDtoPlayerIDStore) GetByID(id int64) (lst []int64, err error) {
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

// Set them 1 tableid
func (trID *TableIDtoPlayerIDStore) Set(id int64) {
	trID.RLock()
	defer trID.RUnlock()
	trID.data[id] = nil
	return
}

// Remove xoa 1 tableid
func (trID *TableIDtoPlayerIDStore) Remove(id int64) {
	trID.RLock()
	defer trID.RUnlock()
	delete(trID.data, id)
	return
}

// SetPlayer them 1 player id vao data
func (trID *TableIDtoPlayerIDStore) SetPlayer(idTable int64, idPlayer int64) {
	trID.RLock()
	defer trID.RUnlock()
	trID.data[idTable] = append(trID.data[idTable], idPlayer)
	return
}

// RemovePlayer xoa 1 player id khoi data
func (trID *TableIDtoPlayerIDStore) RemovePlayer(idTable int64, idPlayer int64) error {
	trID.RLock()
	defer trID.RUnlock()
	for i, v := range trID.data[idTable] {
		if v == idPlayer {
			trID.data[idTable] = append(trID.data[idTable][:i], trID.data[idTable][i+1:]...)
			break
		}
	}
	return errSC.Success
}
