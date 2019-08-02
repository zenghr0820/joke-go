package router

import (
	"github.com/gin-gonic/gin"
	"joke-go/controller"
)

// 初始化路由
func Init() *gin.Engine {
	// 初始化引擎
	r := gin.Default()

	// 注册路由组
	api := r.Group("/joke/v1.0/api")
	{
		//api.GET("/create", controller.Create)
		api.GET(":type", controller.Fetch)
	}

	return r
}
