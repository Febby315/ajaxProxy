package utils

import (
	"log"
	"net/http"

	"github.com/widuu/goini"
)

var config *goini.Config

func init() {
	config = goini.SetConfig("./config.ini") //读取配置文件
}

// GetConfValue 获取字符串配置
func GetConfValue(secName string, keyName string) string {
	return config.GetValue(secName, keyName)
}

// EnableXDA 允许跨域及访问日志
func EnableXDA(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	r.ParseForm()
	log.Printf("--> %s %s", r.URL.String(), r.Method)
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问域r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Credentials", "true")     //允许访问传输Cookies
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") //header的类型
	// w.Header().Add("Access-Control-Allow-Headers", "Accept,Accept-Language,Content-Language,Content-Type") //header的类型
	// w.Header().Set("Access-Control-Request-Method", "POST") //允许访问方式
	// w.Header().Set("Content-Type", "application/json")
	return w
}
