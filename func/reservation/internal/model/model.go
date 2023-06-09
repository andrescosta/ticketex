package model

type Reservation struct {
	ID        string `json:"id"`
	Adventure struct {
		ID string `json:"id"`
	} `json:"adventure"`
	Capacity struct {
		Type    string `json:"type"`
		Current int    `json:"current"`
		Max     int    `json:"max"`
	} `json:"capacity"`
}
