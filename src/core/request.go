package core

import (
	"mime/multipart"
	"net/http"
	"io/ioutil"
	"github.com/wailovet/osmanthuswine/src/session"
	"net/url"
	"log"
)

type Request struct {
	GET     map[string]string
	POST    map[string]string
	REQUEST map[string]string
	COOKIE  map[string]string
	SESSION map[string]string
	HEADER  map[string]string
	BODY    string
	FILES   map[string][]*multipart.FileHeader
	OriginRequest *http.Request
}

func (r *Request) SyncGetData(request *http.Request) {
	r.OriginRequest = request
	get := request.URL.Query()
	r.GET = make(map[string]string)
	for k := range get {
		str := request.URL.Query().Get(k)
		tmp, err := url.QueryUnescape(str)
		if err != nil {
			log.Println(err.Error())
			r.GET[k] = str
			r.REQUEST[k] = str
		}else{
			r.GET[k] = tmp
			r.REQUEST[k] = tmp
		}
	}
}

func (r *Request) SyncPostData(request *http.Request, mem int64) {
	r.OriginRequest = request
	request.ParseForm()
	request.ParseMultipartForm(mem)
	r.POST = make(map[string]string)

	post := request.PostForm
	for k := range post {
		str := request.PostFormValue(k)
		tmp, err := url.QueryUnescape(str)
		if err != nil {
			log.Println(err.Error())
			r.POST[k] = str
			r.REQUEST[k] = str
		}else{
			r.POST[k] = tmp
			r.REQUEST[k] = tmp
		}
	}

	if request.MultipartForm != nil {
		r.FILES = request.MultipartForm.File
		mf := request.MultipartForm.Value
		for k := range mf {
			if len(mf[k]) > 0 {
				r.POST[k] = mf[k][0]
				r.REQUEST[k] = mf[k][0]
			}
		}
	}

	body, _ := ioutil.ReadAll(request.Body)
	r.BODY = string(body)
}

func (r *Request) SyncHeaderData(request *http.Request) {
	r.OriginRequest = request
	r.HEADER = make(map[string]string)
	header := request.Header
	for k := range header {
		if len(header[k]) > 0 {
			r.HEADER[k] = header[k][0]
		}
	}

}

func (r *Request) SyncCookieData(request *http.Request) {
	r.OriginRequest = request
	cookie := request.Cookies()
	r.COOKIE = make(map[string]string)
	for k := range cookie {
		r.COOKIE[cookie[k].Name] = cookie[k].Value
	}
}

func (r *Request) SyncSessionData(session *session.Session) {
	r.SESSION = session.GetSession()
}
