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
	Session   *session.Session
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

func (r *Response) SetSession(name string, value string) {
	data := r.Session.GetSession()
	data[name] = value
	r.Session.SetSession(data)
}

func (r *Response) DeleteSession(name string) {
	data := r.Session.GetSession()
	delete(data, name)
	r.Session.SetSession(data)
}

func (r *Response) ClearSession() {
	data := make(map[string]string)
	r.Session.SetSession(data)
}

func (r *Response) SetCookie(name string, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Secure:   false,
		HttpOnly: false,
	}
	http.SetCookie(r.ResWriter, cookie)
}

func (r *Response) SetHeader(name string, value string) {
	r.ResWriter.Header().Set(name, value)
}
