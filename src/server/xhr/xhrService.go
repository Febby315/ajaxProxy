package xhr

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"../../utils"
)

// Post 代理请求接口
func Post(w http.ResponseWriter, r *http.Request) {
	utils.EnableXDA(w, r)                //跨域请求及日志
	xhr := Xhr{}                         //创建一个XHR对象
	json.NewDecoder(r.Body).Decode(&xhr) //赋值
	//
	client := &http.Client{} // 创建请求客户端
	req, _ := http.NewRequest(xhr.Method, xhr.URL, strings.NewReader(xhr.Body))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36") // 伪装浏览器
	// 设置用户自定义请求头
	for k, v := range xhr.Header {
		req.Header.Set(k, v)
	}
	if res, err := client.Do(req); err == nil {
		log.Printf("<-- %v %v %v\n", xhr.URL, xhr.Method, res.StatusCode)
		defer res.Body.Close()
		// 转发服务器header
		for k, v := range res.Header {
			w.Header().Set(k, v[0])
		}
		body, _ := ioutil.ReadAll(res.Body)
		w.Write(body) // 请求数据并返回给客户端
	} else {
		w.WriteHeader(500)
		w.Write([]byte(err.Error())) // 请求数据并返回给客户端
	}
}

// Get 代理请求静态资源
func Get(w http.ResponseWriter, r *http.Request) {
	utils.EnableXDA(w, r)                // 跨域请求及日志
	xhr := Xhr{}                         // 创建一个XHR对象
	json.NewDecoder(r.Body).Decode(&xhr) // 赋值
	xhr.URL = r.Form.Get("url")
	//
	client := &http.Client{} // 创建请求客户端
	req, _ := http.NewRequest(xhr.Method, xhr.URL, strings.NewReader(xhr.Body))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36") // 伪装浏览器
	// 设置用户自定义请求头
	for k, v := range xhr.Header {
		req.Header.Set(k, v)
	}
	if res, err := client.Do(req); err == nil {
		log.Printf("<-- %v %v %v\n", xhr.URL, xhr.Method, res.StatusCode)
		defer res.Body.Close()
		// 转发服务器header
		for k, v := range res.Header {
			w.Header().Set(k, v[0])
		}
		body, _ := ioutil.ReadAll(res.Body)
		w.Write(body) // 请求数据并返回给客户端
	} else {
		w.WriteHeader(500)
		w.Write([]byte(err.Error())) // 请求数据并返回给客户端
	}
}

// Proxy mux代理
func Proxy(w http.ResponseWriter, r *http.Request) {
	utils.EnableXDA(w, r)                // 跨域请求及日志
	xhr := Xhr{}                         // 创建一个XHR对象
	json.NewDecoder(r.Body).Decode(&xhr) // 赋值
	scheme := mux.Vars(r)["scheme"]
	xhr.URL = strings.Replace(r.URL.String(), "/"+scheme+"/", scheme+"://", -1)
	//
	client := &http.Client{} // 创建请求客户端
	req, _ := http.NewRequest(xhr.Method, xhr.URL, strings.NewReader(xhr.Body))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36") // 伪装浏览器
	// 设置用户自定义请求头
	for k, v := range xhr.Header {
		req.Header.Set(k, v)
	}
	if res, err := client.Do(req); err == nil {
		log.Printf("<-- %v %v %v\n", xhr.URL, xhr.Method, res.StatusCode)
		defer res.Body.Close()
		// 转发服务器header
		for k, v := range res.Header {
			w.Header().Set(k, v[0])
		}
		body, _ := ioutil.ReadAll(res.Body)
		w.Write(body) // 转发请求数据
	} else {
		w.WriteHeader(500)
		w.Write([]byte(err.Error())) // 请求数据并返回给客户端
	}
}
