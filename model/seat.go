package model

type ModifySeat struct {
	HallID uint   `json:"id"`
	Seat   string `json:"seat"`
}
