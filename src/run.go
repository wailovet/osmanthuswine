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
	"encoding/json"
	"./struct"
	"./registered"
	"../config"
	"unicode"
	"reflect"
)

func getModuleName(name string) string {
	if name == "" {
		return "index"
	}
	return strings.ToLower(name)
}
func getControllerName(name string) string {
	if name == "" {
		return "Index"
	}
	for i, v := range name {
		return string(unicode.ToUpper(v)) + name[i+1:]
	}
	return "Index"
}
func getFunName(name string) string {
	if name == "" {
		return "Index"
	}
	for i, v := range name {
		return string(unicode.ToUpper(v)) + name[i+1:]
	}
	return "Index"
}

func Run() {

	path, _ := GetCurrentPath()
	os.Chdir(path)
	log.Println("工作目录:", path)

	config.RegisteredInit()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.HandleFunc("/Api/*", func(writer http.ResponseWriter, request *http.Request) {
		sar := strings.Split(request.URL.Path, "/")

		for ; len(sar) < 5; {
			sar = append(sar, "")
		}

		module := getModuleName(sar[2])
		controller := getControllerName(sar[3])
		fun := getFunName(sar[4])

		ctr := "*" + module + "." + controller
		_, ok := registered.RegisteredData[ctr]
		if ok {
			f := registered.RegisteredData[ctr].MethodByName(fun)
			if f.IsValid() {
				requestData := owstruct.Request{}
				get := request.URL.Query()
				requestData.GET = make(map[string]string)
				for k := range get {
					requestData.GET[k] = request.URL.Query().Get(k)
				}

				body, _ := ioutil.ReadAll(request.Body)
				requestData.BODY = string(body)

				responseHandle := owstruct.Response{ResWriter: writer}
				f.Call([]reflect.Value{reflect.ValueOf(requestData), reflect.ValueOf(responseHandle)})
			} else {
				writer.WriteHeader(404)
			}
		} else {
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

	configText, err := ioutil.ReadFile("./config/main.json")

	cc := new(owstruct.Config)
	if err != nil {
		log.Println("配置文件错误,启动失败:", err.Error())
		os.Exit(0)
	}

	json.Unmarshal(configText, cc)
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
