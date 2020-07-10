package helper

import (
	"archive/zip"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetSelfFilePath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return ""
	}
	path, _ = filepath.Abs(string(path[0 : i+1]))
	return path
}

func CleanSuperfluousSpace(s string) string {
	for strings.Index(s, "  ") > -1 {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return strings.TrimSpace(s)
}

func Md5ToString(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func GetLocalIP() string {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
		err     error
		ipv4    string
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return ""
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
			}
		}
	}
	return ipv4
}

func JsonByFile(file string, v interface{}) {
	data, _ := ioutil.ReadFile(file)
	json.Unmarshal(data, v)
}

func JsonToFile(file string, v interface{}) bool {
	data, err := json.Marshal(v)
	if err != nil {
		return false
	}
	return ioutil.WriteFile(file, data, 0644) == nil
}

func JsonEncode(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func GetFreePort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

func Interface2Struct(in interface{}, out interface{}) {
	raw, _ := json.Marshal(in)
	json.Unmarshal(raw, &out)
}
func Interface2Map(in interface{}, out map[string]interface{}) {

	raw, _ := json.Marshal(in)
	json.Unmarshal(raw, &out)
}

func Interface2Interface(in interface{}) (out interface{}) {
	raw, _ := json.Marshal(in)
	json.Unmarshal(raw, &out)
	return
}

func CleanExtraCharacters(a string, b string) string {
	for strings.Index(a, b+b) > -1 {
		a = strings.Replace(a, b+b, b, -1)
	}
	return a
}

var isUseConsul = false

func IsUseConsul() bool {
	return isUseConsul
}
func UseConsul() {
	isUseConsul = true
}

var consulClientId string

func SetConsulClientId(id string) {
	consulClientId = id
}

func GetConsulClientId() string {
	return consulClientId
}

func InArray(t string, arr []string) bool {

	for e := range arr {
		if arr[e] == t {
			return true
		}
	}
	return false
}

func Unzip(file string, path string) error {
	File, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer File.Close()
	for _, v := range File.File {
		info := v.FileInfo()
		fileName, _ := filepath.Abs(path + "/" + v.Name)
		_ = os.RemoveAll(fileName)
		if info.IsDir() {
			err := os.MkdirAll(fileName, 0777)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		srcFile, err := v.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer srcFile.Close()

		newFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			continue
		}
		io.Copy(newFile, srcFile)
		newFile.Close()
	}
	return nil
}

func RandomInt(length int) int {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte

	for i := 0; i < length; {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i) + int64(time.Now().Nanosecond())))
		b := bytes[r.Intn(len(bytes))]
		if i == 0 && b == '0' {
			continue
		}
		result = append(result, b)
		i++
	}
	num, _ := strconv.Atoi(string(result))
	return num
}

func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < length; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i) + int64(time.Now().Nanosecond())))
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
