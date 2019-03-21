package main

import (
	"log"
	"net/http"

	"./server/xhr"
	"./utils"
	"github.com/julienschmidt/httprouter"
)

//声明全局变量
var R = httprouter.New()

//初始化
func init() {
	R.POST("/api/post", xhr.Post)
	R.GET("/api/get", xhr.Get)
}

//程序主入口
func main() {
	//启动服务
	port := utils.GetConfValue("global", "port")
	if err := http.ListenAndServe(":"+port, R); err != nil {
		log.Printf("服务启动失败:\n%s", err.Error())
	}
}
