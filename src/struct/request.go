package owstruct

import (
	"mime/multipart"
	"net/http"
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

func (r *Request) SyncPostData(request *http.Request,mem int64) {
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
}
