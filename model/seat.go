package model

type ModifySeat struct {
	HallID int    `json:"id"`
	Seat   string `json:"seat"`
}
