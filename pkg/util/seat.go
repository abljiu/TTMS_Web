package util

import (
	"TTMS_Web/model"
	"strconv"
	"strings"
)

func ParseSeat(seat string) (seats []int) {
	str := strings.Split(seat, ",")
	for _, numStr := range str {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return
		}
		// 将整数添加到切片中
		seats = append(seats, int(num))
	}
	return
}

func UpdateSessionSeat(session *model.Session, seat string, num int) {
	seats := ParseSeat(seat)
	row := session.SeatRow
	bytes := []byte(session.SeatStatus)
	for i, j := 0, 1; j < len(seats); i, j = i+1, j+1 {
		bytes[(seats[i]-1)*row+seats[j]-1] = 2
	}
	session.SurplusTicket -= num
	session.SeatStatus = string(bytes)
}
