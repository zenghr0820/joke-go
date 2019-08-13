package router

import (
	"github.com/gin-gonic/gin"
	"joke-go/controller"
)

// 初始化路由
func Init() *gin.Engine {
	// 初始化引擎
	r := gin.Default()

	// 自定义 gin 日志打印输出
	//r := gin.New()
	//r.Use(middleware.GinLogger())

	// 注册路由组
	api := r.Group("/joke/api")
	{
		//api.GET("/create", controller.Create)
		api.GET("/fetch", controller.Fetch)
		api.GET("/count", controller.Count)
		api.POST("/pull", controller.LoadRefresh)
		api.GET("/get/:id", controller.FindJokeInfo)
	}

	return r
}
