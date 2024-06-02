package serializer

import (
	"TTMS_Web/model"
)

type Theater struct {
	Name    string `json:"name"`
	ID      uint   `json:"id"`
	Address string `json:"address"`
	HallNum int    `json:"hall_num"`
}

func BuildTheater(theater *model.Theater) *Theater {
	return &Theater{
		Name:    theater.Name,
		ID:      theater.ID,
		Address: theater.Address,
		HallNum: theater.HallNum,
	}
}
func BuildTheaters(items []*model.Theater) (Theaters []*Theater) {
	for _, i := range items {
		theater := BuildTheater(i)
		Theaters = append(Theaters, theater)
	}
	return Theaters
}
