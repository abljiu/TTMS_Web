package serializer

import (
	"TTMS_Web/model"
	"strconv"
	"strings"
)

type Hall struct {
	ID         uint
	Name       string
	TheaterID  uint
	SeatRow    int
	SeatColumn int
	Seat       [][]int
	SeatNum    int
}

func BuildHall(item *model.Hall) Hall {
	seatArr := strings.Split(item.Seat, ",")
	var seat [][]int
	seat = make([][]int, item.SeatRow)
	for i := range seat {
		seat[i] = make([]int, item.SeatColumn)
	}
	//fmt.Println(item.ID, item.SeatRow, item.SeatColumn, seatArr, len(seatArr))
	k := 0
	for i := 0; i < item.SeatRow; i++ {
		for j := 0; j < item.SeatColumn; j++ {
			seat[i][j], _ = strconv.Atoi(seatArr[k])
			k++
		}
	}

	return Hall{
		ID:         item.ID,
		Name:       item.Name,
		TheaterID:  item.TheaterID,
		SeatRow:    item.SeatRow,
		SeatColumn: item.SeatColumn,
		Seat:       seat,
		SeatNum:    item.SeatNum,
	}
}

func BuildHalls(items []*model.Hall) (products []Hall) {
	for i := 0; i < len(items); i++ {
		product := BuildHall(items[i])
		products = append(products, product)
	}
	return products
}
