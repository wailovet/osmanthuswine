package owstruct

import "mime/multipart"

type Request struct {
	GET     map[string]string
	POST    map[string]string
	REQUEST map[string]string
	COOKIE  map[string]string
	SESSION map[string]string
	HEADER  map[string]string
	BODY    string
	FILES    map[string][]*multipart.FileHeader
}
