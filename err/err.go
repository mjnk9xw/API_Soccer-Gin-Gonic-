package errSC

import (
	"errors"
)

var (
	ErrIDInt       = errors.New("ID not Int")
	ErrTourName    = errors.New("Da co Tour")
	ErrIDTour      = errors.New("not Tour")
	ErrIDRound     = errors.New("not Round")
	ErrIDTable     = errors.New("not Table")
	ErrIDPlayer    = errors.New("not Player")
	ErrPlayerName  = errors.New("Da co Player")
	ErrIDMatch     = errors.New("not Match")
	DataPost       = errors.New("Data POST err")
	Success        = errors.New("Success")
	NotFound       = errors.New("Not found")
	MapDescription = map[error]string{
		ErrIDInt:      "ID khong phai kieu Int",
		ErrTourName:   "Tour da ton tai",
		ErrIDTour:     "Tour chua co",
		ErrIDRound:    "Round chua co",
		ErrIDTable:    "Table chua co",
		ErrIDPlayer:   "Player chua co",
		ErrPlayerName: "Doi bong da ton tai",
		ErrIDMatch:    "Match chua co",
		DataPost:      "Data Post khong dung dinh dang",
		Success:       "Thanh cong",
		NotFound:      "Khong tim thay",
	}
	MapErrorCode = map[error]int{
		ErrIDInt:      400,
		ErrIDTour:     401,
		ErrTourName:   402,
		ErrIDRound:    403,
		ErrIDTable:    405,
		ErrIDPlayer:   406,
		ErrPlayerName: 407,
		ErrIDMatch:    408,
		DataPost:      500,
		Success:       200,
		NotFound:      404,
	}
)
