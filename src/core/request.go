package osmanthuswine

import (
	"mime/multipart"
	"net/http"
	"io/ioutil"
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
}

func (r *Request) SyncGetData(request *http.Request) {
	get := request.URL.Query()
	r.GET = make(map[string]string)
	for k := range get {
		r.GET[k] = request.URL.Query().Get(k)
	}
}

func (r *Request) SyncPostData(request *http.Request, mem int64) {
	request.ParseForm()
	request.ParseMultipartForm(mem)
	r.POST = make(map[string]string)

	post := request.PostForm
	for k := range post {
		r.POST[k] = request.PostFormValue(k)
	}

	if request.MultipartForm != nil {
		r.FILES = request.MultipartForm.File
		mf := request.MultipartForm.Value
		for k := range mf {
			if len(mf[k]) > 0 {
				r.POST[k] = mf[k][0]
			}
		}
	}

	body, _ := ioutil.ReadAll(request.Body)
	r.BODY = string(body)
}

func (r *Request) SyncHeaderData(request *http.Request) {
	r.HEADER = make(map[string]string)
	header := request.Header
	for k := range header {
		if len(header[k]) > 0 {
			r.HEADER[k] = header[k][0]
		}
	}

}

func (r *Request) SyncCookieData(request *http.Request) {
	cookie := request.Cookies()
	r.COOKIE = make(map[string]string)
	for k := range cookie {
		r.COOKIE[cookie[k].Name] = cookie[k].Value
	}
}
