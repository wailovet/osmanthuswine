package core

import (
	"net/http"
	"encoding/json"
	"github.com/wailovet/osmanthuswine/src/session"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Response struct {
	ResWriter http.ResponseWriter
	session   session.Session
}

func (r *Response) DisplayByString(data string) {
	r.ResWriter.Write([]byte(data))
}
func (r *Response) DisplayByRaw(data []byte) {
	r.ResWriter.Write(data)
}

func (r *Response) Display(data interface{}, msg string, code int) {
	result := ResponseData{code, data, msg}
	text, err := json.Marshal(result)
	if err != nil {
		r.ResWriter.WriteHeader(500)
		r.DisplayByString("服务器异常:" + err.Error())
	}
	r.DisplayByRaw(text)
}

func (r *Response) DisplayByError(msg string, code int) {
	result := ResponseData{code, nil, msg}
	text, err := json.Marshal(result)
	if err != nil {
		r.Display(nil, "JSON返回格式解析异常:"+err.Error(), 500)
	}
	r.DisplayByRaw(text)
}

func (r *Response) DisplayBySuccess(msg string) {
	result := ResponseData{0, nil, msg}
	text, err := json.Marshal(result)
	if err != nil {
		r.Display(nil, "JSON返回格式解析异常:"+err.Error(), 500)
	}
	r.DisplayByRaw(text)
}

func (r *Response) DisplayByData(data interface{}) {
	result := ResponseData{0, data, ""}
	text, err := json.Marshal(result)
	if err != nil {
		r.Display(nil, "JSON返回格式解析异常:"+err.Error(), 500)
	}
	r.DisplayByRaw(text)
}

func (r *Response) SetSession(value map[string]string) {
	r.session.SetSession(value)
}

func (r *Response) SetSessionValue(name string, value string) {
	data := r.session.GetSession()
	data[name] = value
	r.session.SetSession(data)
}
