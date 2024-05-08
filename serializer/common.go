package serializer

import (
	"TTMS_Web/pkg/e"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: http.StatusOK,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: e.GetMsg(http.StatusOK),
	}
}
