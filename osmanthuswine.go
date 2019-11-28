package osmanthuswine

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/wailovet/osmanthuswine/src/core"
	"github.com/wailovet/osmanthuswine/src/helper"
	"github.com/wailovet/osmanthuswine/src/session"
	"github.com/wailovet/overseer"
	"github.com/wailovet/overseer/fetcher"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

var chiRouter *chi.Mux

func GetChiRouter() *chi.Mux {
	if chiRouter == nil {
		chiRouter = chi.NewRouter()
		chiRouter.Use(middleware.RequestID)
		chiRouter.Use(middleware.RealIP)
		chiRouter.Use(middleware.Logger)
		chiRouter.Use(middleware.Recoverer)
	}
	return chiRouter
}
func Run() {
	path, _ := GetCurrentPath()
	os.Chdir(path)
	log.Println("工作目录:", path)
	cc := core.GetInstanceConfig()

	if runtime.GOOS == "windows" || cc.UpdatePath == "" {
		listener, err := net.Listen("tcp", cc.Host+":"+cc.Port)
		if err != nil {
			log.Fatal(err.Error())
		}
		RunProg(overseer.State{
			Listener: listener,
		})
	} else {
		overseer.Run(overseer.Config{
			Program: RunProg,
			Address: cc.Host + ":" + cc.Port,
			Fetcher: &fetcher.File{
				Path:     cc.UpdateDir + cc.UpdatePath,
				Interval: time.Second * 10,
			},
		})
	}

}
func RunProg(state overseer.State) {

	cc := core.GetInstanceConfig()
	helper.GetInstanceLog().Out("开始监听:", cc.Host+":"+cc.Port)

	r := GetChiRouter()

	apiRouter := cc.ApiRouter

	if apiRouter == "" {
		apiRouter = "/Api/*"
	}

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

		responseHandle := core.Response{OriginResponseWriter: writer, Session: sessionMan}

		defer func() {
			defer func() {
				recover()
			}()

			errs := recover()
			if errs == nil {
				return
			}
			errtxt := fmt.Sprintf("%v", errs)
			if errtxt != "" {
				responseHandle.DisplayByError(errtxt, 500, strings.Split(string(debug.Stack()), "\n\t")...)
			}
		}()

		core.GetInstanceRouterManage().RouterSend(request.URL.Path, requestData, responseHandle, cc.CrossDomain)

	})
	if cc.StaticRouter == "" {
		cc.StaticRouter = "/*"
	}
	_, err := os.Stat("./html")
	if err == nil {
		//兼容旧版本
		r.HandleFunc(cc.StaticRouter, func(writer http.ResponseWriter, request *http.Request) {
			path := request.URL.Path
			if path == "/" {
				path = "/index.html"
			}

			helper.GetInstanceLog().Out("静态文件:", "./html"+path)

			f, err := os.Stat("./html" + path)
			if err == nil {
				if f.IsDir() {
					path += "/index.html"
				}
				data, err := ioutil.ReadFile("./html" + path)
				if err == nil {
					writer.Write(data)
					return
				}
			}
			writer.WriteHeader(404)
			writer.Write([]byte(err.Error()))
		})
	} else {
		if cc.StaticFileSystem == nil {
			cc.StaticFileSystem = http.Dir("static")
		}
		r.Handle(cc.StaticRouter, http.FileServer(cc.StaticFileSystem))
	}

	InstanceListener = state.Listener
	http.Serve(state.Listener, r)
	//http.ListenAndServe(cc.Host+":"+cc.Port, r)

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
	return string(path[0 : i+1]), nil
}

var InstanceListener net.Listener

func HandleFunc(pattern string, callback func(request core.Request, response core.Response)) {

	cc := core.GetInstanceConfig()
	r := GetChiRouter()
	r.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {

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

		responseHandle := core.Response{OriginResponseWriter: writer, Session: sessionMan}

		defer func() {
			errs := recover()
			if errs == nil {
				return
			}
			errtxt := fmt.Sprintf("%v", errs)
			if errtxt != "" {
				responseHandle.DisplayByError(errtxt, 500, strings.Split(string(debug.Stack()), "\n\t")...)
			}
		}()

		responseHandle.OriginResponseWriter.Header().Set("Content-Type", "application/json;charset=UTF-8")
		if cc.CrossDomain != "" {
			responseHandle.OriginResponseWriter.Header().Set("Access-Control-Allow-Origin", cc.CrossDomain)
		}

		callback(requestData, responseHandle)
	})
}
