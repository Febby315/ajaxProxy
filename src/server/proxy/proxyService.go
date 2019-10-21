package proxy

import (
	"log"
	"net/http"

	"../../utils"
)

// Get 代理请求静态资源
func Get(w http.ResponseWriter, r *http.Request) {
	utils.EnableXDA(w, r) // 跨域请求及日志
	log.Printf("<-> %s %s\n", r.Method, r.Form.Get("proxy"))
	if proxy := r.Form.Get("proxy"); proxy != "" {
		req, _ := http.NewRequest(r.Method, proxy, r.Body)
		req.Header = r.Header
		req.Form = r.Form
		utils.SendRequest(w, req)
	}
}

// Post 代理请求接口
func Post(w http.ResponseWriter, r *http.Request) {
	utils.EnableXDA(w, r) // 跨域请求及日志
	log.Printf("<-> %s %s\n", r.Method, r.Header.Get("Proxy"))
	if proxy := r.Header.Get("Proxy"); proxy != "" {
		req, _ := http.NewRequest(r.Method, proxy, r.Body)
		req.Header = r.Header
		req.Form = r.Form
		utils.SendRequest(w, req)
	}
}
