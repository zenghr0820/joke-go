package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"joke-go/logger"
	"joke-go/models"
	"strconv"
	"strings"
	"time"
)

var types = map[string]bool{
	"text":  true,
	"pic":   true,
	"video": true,
	"hot":   true,
}

// 用于获取糗事分类数据
func Fetch(ctx *gin.Context) {

	var (
		err         error
		jokes       models.Jokes
		jokeType    string
		limit       int
		currentPage int
		start       int
	)

	jokeType = ctx.Param("type")

	logger.Info("===> 请求参数：", jokeType)

	// 判断type是否属于四种类型中的一种
	if _, ok := types[jokeType]; !ok {
		// 不存在
		logger.Info("===> aaaaa：", jokeType)
		ctx.JSON(200, gin.H{
			"code": 40002,
			"msg":  "Invalid request params!",
		})
		return
	}

	// 检查是否有分页信息
	limitS := ctx.Param("limit")
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 40002,
				"msg":  "Invalid request params!",
			})
			return
		}
		if limit <= 0 {
			limit = 16
		}
	} else {
		limit = 16
	}

	// 查看分页
	currentPageS := ctx.Param("page")
	if currentPageS != "" {
		currentPage, err = strconv.Atoi(limitS)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 40002,
				"msg":  "Invalid request params!",
			})
			return
		}
		if currentPage <= 0 {
			currentPage = 1
		}
	} else {
		currentPage = 1
	}

	// 开始位置 和 结束位置
	start = (currentPage - 1) * limit
	//stop = start + limit

	// 查询数据
	switch jokeType {
	case "hot":
		err = models.Orm.Order("time desc").Offset(start).Limit(limit).Find(&jokes).Error
	case "pic":
		err = models.Orm.Where("image_url <> ''").Or("video_url <> ''").Offset(start).Limit(limit).Find(&jokes).Error
	case "video":
		err = models.Orm.Where("video_url <> ''").Or("image_url <> ''").Offset(start).Limit(limit).Find(&jokes).Error
	case "text":
		err = models.Orm.Where("image_url = '' AND video_url = ''").Offset(start).Limit(limit).Find(&jokes).Error
	}

	if err != nil {
		logger.Error(err)
		ctx.JSON(500, gin.H{
			"code": 50001,
			"msg":  "connect error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 0,
		"data": jokes,
		"msg":  "success",
	})
}

func Demo(ctx *gin.Context) {
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(convertLogLevel("info"))

	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func Create(ctx *gin.Context) {

	var (
		err  error
		joke models.Joke
	)

	joke.Id = "2018"
	joke.Type = "hot"
	joke.Title = "cehsi"
	joke.Time = time.Now()

	err = models.Orm.Create(&joke).Error

	if err != nil {
		logger.Error(err)
		ctx.JSON(500, gin.H{
			"code": 50001,
			"msg":  "connect error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 0,
		"data": joke,
		"msg":  "success",
	})

}

// 把字符串转换为日志级别（数字）
func convertLogLevel(levelStr string) zapcore.Level {
	// 不区分大小写
	levelStr = strings.ToLower(levelStr)
	var level zapcore.Level
	switch levelStr {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	return level
}
