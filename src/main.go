package main

import (
	"net/http"

	"./server/xhr"
	"./utils"
	"github.com/julienschmidt/httprouter"
)

//声明全局变量
var R = httprouter.New()

//初始化
func init() {
	R.POST("/api/xhr", xhr.Post)
}

//程序主入口
func main() {
	//启动服务
	port := utils.GetConfValue("global", "port")
	http.ListenAndServe(":"+port, R)
}
