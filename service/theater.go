package service

type TheaterService struct {
	TheaterID   uint   `form:"theater_id" json:"theater_id"`
	TheaterName string `form:"theater_name" json:"theater_name"`
	SeatNum     uint   `form:"seat_num" json:"seat_num"`
	AddressID   uint   `form:"address_id" json:"address_id"`
}

//// Add 添加剧院
//func (service *TheaterService) Add(ctx context.Context) serializer.Response {
//	code := e.Success
//
//	theaterDao := dao.NewTheaterDao(ctx)
//	addressDao:=dao.
