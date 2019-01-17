package osmanthuswine

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"time"
	"net/http"
	"io/ioutil"
	"os/exec"
	"os"
	"path/filepath"
	"strings"
	"errors"
	"log"
	"reflect"
	"github.com/wailovet/osmanthuswine/src/core"
)

func Run() {
	path, _ := GetCurrentPath()
	os.Chdir(path)
	log.Println("工作目录:", path)

	cc := core.Config{}
	cc.ReadConfig("./config/main.json")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.HandleFunc("/Api/*", func(writer http.ResponseWriter, request *http.Request) {

		requestData := core.Request{}

		//GET
		requestData.SyncGetData(request)
		//POST
		requestData.SyncPostData(request, cc.PostMaxMemory)
		//HEADER
		requestData.SyncHeaderData(request)
		//COOKIE
		requestData.SyncCookieData(request)

		responseHandle := core.Response{ResWriter: writer}

		ok := core.GetInstanceRouterManage().RouterSend(request.URL.Path, requestData, responseHandle)
		if ok == nil {
			writer.WriteHeader(404)
		}
		writer.Write([]byte(""))

	})

	r.HandleFunc("/*", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if path == "/" {
			path = "/index.html"
		}
		println("静态文件:", "./html"+path)
		data, err := ioutil.ReadFile("./html" + path)
		if err == nil {
			writer.Write([]byte(data))
		} else {
			writer.WriteHeader(404)
			writer.Write([]byte(""))
		}
	})

	log.Println("开始监听:", cc.Host+":"+cc.Port)
	http.ListenAndServe(cc.Host+":"+cc.Port, r)
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0: i+1]), nil
}
