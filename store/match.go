package store

import (
	"Soccer-oauth2/err"
	"Soccer-oauth2/interfacemodel"
	"sync"
)

// MatchStore data Match
type MatchStore struct {
	sync.RWMutex
	data map[int64]interfacemodel.MatchInfo
}

// NewMatchStore khoi tao data Match
func NewMatchStore() *MatchStore {
	return &MatchStore{
		data: make(map[int64]interfacemodel.MatchInfo),
	}
}

// GetByID lay 1 Match theo id trong data Match
func (ts *MatchStore) GetByID(id int64) (tf interfacemodel.MatchInfo, err error) {
	ts.RLock()
	defer ts.RUnlock()
	if t, ok := ts.data[id]; ok {
		err = errSC.Success
		tf = t
		return
	}
	err = errSC.ErrIDMatch
	return

}

// Set tao key value trong map data Match
func (ts *MatchStore) Set(tf interfacemodel.MatchInfo) {
	ts.RLock()
	defer ts.RUnlock()
	ts.data[tf.GetMatchID()] = tf
	return
}

// Remove xoa 1 key trong data Match
func (ts *MatchStore) Remove(id int64) error {
	ts.RLock()
	defer ts.RUnlock()
	_, err := ts.data[id]
	if err {
		delete(ts.data, id)
		return errSC.Success
	}
	return errSC.NotFound
}
