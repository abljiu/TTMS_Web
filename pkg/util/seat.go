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
	for i, j := 0, 1; j < len(seats); i, j = i+2, j+2 {
		bytes[(seats[i]-1)*row*2+2*seats[j]-2] = '2'
	}
	session.SurplusTicket -= num
	session.SeatStatus = string(bytes)
}

func ReturnSessionSeat(session *model.Session, seat string, num int) {
	seats := ParseSeat(seat)
	row := session.SeatRow
	bytes := []byte(session.SeatStatus)
	for i, j := 0, 1; j < len(seats); i, j = i+2, j+2 {
		bytes[(seats[i]-1)*row*2+2*seats[j]-2] = '1'
	}
	session.SurplusTicket -= num
	session.SeatStatus = string(bytes)
}

func IsRepeatSeat(seat string, seatStatus string, row int) bool {
	seats := ParseSeat(seat)
	bytes := []byte(seatStatus)
	for i, j := 0, 1; j < len(seats); i, j = i+2, j+2 {
		if bytes[(seats[i]-1)*row*2+2*seats[j]-2] == '2' {
			return true
		}
	}
	return false
}
