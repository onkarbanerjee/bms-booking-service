package models

type Booking struct {
	ID       string `json:"id,omitempty"`
	UserID   string `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	ShowID   int    `json:"show_id,omitempty"`
	CinemaID int    `json:"cinema_id,omitempty"`
	Seats    []Seat `json:"seats,omitempty"`
}
