package smpkg

import (
	"encoding/json"
	"net/http"
)

type responseData struct {
	RequestId string      `json:"request_id"`
	Msg       string      `json:"msg"`
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
}

func SuccessRes(ctx CtxContext, w http.ResponseWriter, data interface{}) {
	resp := responseData{
		RequestId: ctx.RequestID(),
		Code:      200,
		Msg:       "success",
		Data:      data,
	}
	jsonResponse(w, resp)
	return
}

func ErrorRes(ctx CtxContext, w http.ResponseWriter, err error) {
	resp := responseData{
		RequestId: ctx.RequestID(),
		Code:      500,
		Msg:       err.Error(),
		Data:      map[string]interface{}{},
	}
	jsonResponse(w, resp)
	return
}

func jsonResponse(w http.ResponseWriter, resp responseData) {
	body, _ := json.Marshal(resp)
	// 先得调用header才能调用write或者writeString
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// 写入响应实际内容
	_, _ = w.Write(body)
}
