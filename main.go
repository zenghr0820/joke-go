package main

import (
	"joke-go/config"
	"joke-go/logger"
	"joke-go/models"
	"joke-go/router"
	"net/http"
)

func main() {
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	//logger.SetSugar()
	logger.InitLog()
	// 初始化 ORM
	models.InitOrm()
	// 初始化路由
	router := router.Init()

	port := config.GetConfig("app.port")

	logger.Info("监听端口：", port)
	logger.Error("监听端口：")

	s := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	println("[JOKE-GO] Starting...")
	println(" - URL:  	127.0.0.1:", port)

	// 监听服务
	s.ListenAndServe()
}
