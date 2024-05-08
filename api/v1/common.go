package v1

import (
	"TTMS_Web/pkg/e"
	"TTMS_Web/serializer"
	"encoding/json"
	"errors"
)

func ErrorResponse(err error) serializer.Response {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return serializer.Response{
			Status: e.InvalidParams,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: e.InvalidParams,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
