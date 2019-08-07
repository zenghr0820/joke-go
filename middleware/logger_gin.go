package middleware

import (
	"github.com/gin-gonic/gin"
	"joke-go/logger"
	"time"
)

// 自定义 gin 日志打印输出 --- 中间件
func GinLogger() gin.HandlerFunc  {
	return func (c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logger.Info(statusCode, latency, clientIP, method, path)
	}

}
