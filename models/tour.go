package models

import "fmt"

type TourTime struct {
	StartTime string `json:"start"`
}

type Tour struct {
	TourID      int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"des"`
	Time        *TourTime `json:"time"`
}

func (t *Tour) GenIDKey() string {
	str := fmt.Sprintf("%d", t.TourID)
	return str
}
func (t *Tour) GetTourID() int64 {
	return t.TourID
}
func (t *Tour) SetTourID() {
	t.TourID = genNextTourID()
}
func (t *Tour) GetTourName() string {
	return t.Name
}
