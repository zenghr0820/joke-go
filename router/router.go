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
	r.GET("/joke/Demo", controller.Demo)
	api := r.Group("/joke/v1.0/api")
	{
		//api.GET("/create", controller.Create)
		api.GET(":type", controller.Fetch)
	}

	return r
}
