package main

import (
	"joke-go/router"
	"joke-go/utils"
	"net/http"
)

func main() {
	// 初始化日志
	utils.InitLogging()
	// 初始化配置
	utils.InitConfig()
	// 初始化 ORM
	utils.InitOrm()

	router := router.Init()

	port := utils.GetConfig("app.port")

	utils.Log.Info("监听端口：", port)

	s := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// 监听服务
	s.ListenAndServe()
}
