package osmanthuswine

import (
	"github.com/go-chi/chi"
	"net/http"
	"io/ioutil"
	"os/exec"
	"os"
	"path/filepath"
	"strings"
	"errors"
	"log"
	"github.com/wailovet/osmanthuswine/src/core"
	"github.com/wailovet/osmanthuswine/src/helper"
	"github.com/wailovet/osmanthuswine/src/session"
	"github.com/go-chi/chi/middleware"
	"time"
	"fmt"
)

var chiRouter *chi.Mux

func GetChiRouter() *chi.Mux {
	if chiRouter == nil {

		chiRouter = chi.NewRouter()
		chiRouter.Use(middleware.RequestID)
		chiRouter.Use(middleware.RealIP)
		chiRouter.Use(middleware.Logger)
		chiRouter.Use(middleware.Recoverer)
		chiRouter.Use(middleware.Timeout(60 * time.Second))
	}
	return chiRouter
}

func Run() {
	path, _ := GetCurrentPath()
	os.Chdir(path)
	log.Println("工作目录:", path)

	cc := core.GetInstanceConfig()
	r := GetChiRouter()

	apiRouter := cc.ApiRouter

	r.HandleFunc(apiRouter, func(writer http.ResponseWriter, request *http.Request) {

		requestData := core.Request{}

		sessionMan := session.New(request, writer)

		requestData.REQUEST = make(map[string]string)
		//GET
		requestData.SyncGetData(request)
		//POST
		requestData.SyncPostData(request, cc.PostMaxMemory)
		//HEADER
		requestData.SyncHeaderData(request)
		//COOKIE
		requestData.SyncCookieData(request)
		//SESSION
		requestData.SyncSessionData(sessionMan)

		responseHandle := core.Response{ResWriter: writer, Session: sessionMan}

		defer func() {
			errs := recover()
			if errs == nil {
				return
			}
			responseHandle.DisplayByError(fmt.Sprintf("%v", errs), 500)

		}()

		responseHandle.ResWriter.Header().Set("Content-Type", "application/json;charset=UTF-8")
		if cc.CrossDomain != "" {
			responseHandle.ResWriter.Header().Set("Access-Control-Allow-Origin", cc.CrossDomain)
		}

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

		helper.GetInstanceLog().Out("静态文件:", "./html"+path)
		data, err := ioutil.ReadFile("./html" + path)
		if err == nil {
			writer.Write([]byte(data))
		} else {
			writer.WriteHeader(404)
			writer.Write([]byte(""))
		}
	})
	helper.GetInstanceLog().Out("开始监听:", cc.Host+":"+cc.Port)
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
