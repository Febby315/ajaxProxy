package main

import (
	"log"
	"net/http"

	"./server/proxy"
	"./utils"
	"github.com/gorilla/mux"
)

var port = utils.GetConfValue("global", "port")

// R mux路由容器
var R = mux.NewRouter()

//初始化
func init() {
	R.HandleFunc("/api/get", proxy.Get)
	R.HandleFunc("/api/post", proxy.Post)
}

//程序主入口
func main() {
	//启动服务
	log.Fatalln(http.ListenAndServe("0.0.0.0:"+port, R))
}
