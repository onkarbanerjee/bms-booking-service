package models

import "time"

type Seat struct {
	Row    string  `json:"row,omitempty"`
	Column int     `json:"column,omitempty"`
	Price  float32 `json:"price,omitempty"`
}

type Show struct {
	ID       int       `json:"id,omitempty"`
	CinemaID int       `json:"cinema_id,omitempty"`
	Time     time.Time `json:"time,omitempty"`
	Seats    []Seat    `json:"seats,omitempty"`
}
