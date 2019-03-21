package utils

import (
	"log"
	"net/http"

	"github.com/widuu/goini"
)

var config *goini.Config

func init() {
	//读取配置文件
	config = goini.SetConfig("./config.ini")
}

// 获取字符串配置
func GetConfValue(secName string, keyName string) string {
	return config.GetValue(secName, keyName)
}

//允许跨域及访问日志
func EnableXDA(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	log.Printf("--> %s %s", r.Method, r.URL.String())
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域r.Header.Get("Origin")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Access-Control-Allow-Credentials", "true")     //允许访问所有域
	w.Header().Set("content-type", "application/json")
	return w
}
