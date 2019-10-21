package utils

import (
	"io/ioutil"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")         //允许访问域r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Headers", "*")        //header的类型
	w.Header().Set("Access-Control-Allow-Credentials", "true") //允许访问传输Cookies
	return w
}

// SendRequest 发送请求
func SendRequest(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	r.Header.Set("User-Agent", GetConfValue("global", "User-Agent")) // 伪装浏览器
	client := &http.Client{}                                         // 创建请求客户端
	if res, err := client.Do(r); err == nil {
		defer res.Body.Close()
		// 转发服务器header
		for k, vs := range res.Header {
			for _, v := range vs {
				w.Header().Set(k, v)
			}
		}
		body, _ := ioutil.ReadAll(res.Body)
		w.Write(body) // 请求数据并返回给客户端
		log.Printf("<-- %v %v %v\n", r.URL.String(), r.Method, res.StatusCode)
	} else {
		w.WriteHeader(500)
		w.Write([]byte(err.Error())) // 请求数据并返回给客户端
	}
	return w
}
