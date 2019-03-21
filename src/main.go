package main

import (
	"log"
	"net/http"

	"./server/xhr"
	"./utils"
	"github.com/gorilla/mux"
)

var port = utils.GetConfValue("global", "port")

// R mux路由容器
var R = mux.NewRouter()

//初始化
func init() {
	R.HandleFunc("/api/post", xhr.Post)
	R.HandleFunc("/api/get", xhr.Get)
	R.HandleFunc("/{scheme:(?:http|https)}/{host}/{path:.*}", xhr.Proxy)
}

//程序主入口
func main() {
	//启动服务
	if err := http.ListenAndServe(":"+port, R); err != nil {
		log.Printf("服务启动失败:\n%s", err.Error())
	}
}
