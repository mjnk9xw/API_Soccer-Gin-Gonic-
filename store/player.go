package store

import (
	"Soccer-oauth2/err"
	"Soccer-oauth2/interfacemodel"
	"sync"
)

// PlayerStore data Player
type PlayerStore struct {
	sync.RWMutex
	data map[int64]interfacemodel.PlayerInfo
}

// NewPlayerStore khoi tao data Player
func NewPlayerStore() *PlayerStore {
	return &PlayerStore{
		data: make(map[int64]interfacemodel.PlayerInfo),
	}
}

// GetByID lay Player theo id trong data
func (ts *PlayerStore) GetByID(id int64) (tf interfacemodel.PlayerInfo, err error) {
	ts.RLock()
	defer ts.RUnlock()
	if t, ok := ts.data[id]; ok {
		tf = t
		err = errSC.Success
		return
	}
	err = errSC.ErrIDPlayer
	return

}

// GetName lay tat ca name cua player
func (ts *PlayerStore) GetName() (lst []string, err error) {
	ts.RLock()
	defer ts.RUnlock()
	err = errSC.Success
	lst = make([]string, 0)
	for _, v := range ts.data {
		lst = append(lst, v.GetPlayerName())
	}
	return
}

// Set them player vao data
func (ts *PlayerStore) Set(tf interfacemodel.PlayerInfo) {
	ts.RLock()
	defer ts.RUnlock()
	ts.data[tf.GetPlayerID()] = tf
	return
}

// Remove xoa 1 player trong data
func (ts *PlayerStore) Remove(id int64) error {
	ts.RLock()
	defer ts.RUnlock()
	_, err := ts.data[id]
	if err {
		delete(ts.data, id)
		return errSC.Success
	}
	return errSC.NotFound
}
