package controller

import (
	"github.com/gin-gonic/gin"
	"joke-go/logger"
	"joke-go/models"
	"math/rand"
	"strconv"
	"time"
)

var types = map[string]bool{
	"text":  true,
	"pic":   true,
	"video": true,
	"hot":   true,
}

type PostLoadRefresh struct {
	Time    string   `json:"time" form:"time"`
	JokeIds []string `json:"joke_ids" form:"joke_ids"`
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

	jokeType = ctx.Query("type")

	logger.Info("===> 请求参数：", jokeType)

	// 判断type是否属于四种类型中的一种
	if _, ok := types[jokeType]; !ok {
		// 不存在
		ctx.JSON(200, gin.H{
			"code":    40002,
			"message": "Invalid request params!",
		})
		return
	}

	// 检查是否有分页信息
	limitS := ctx.Query("limit")
	if limitS != "" {
		limit, err = strconv.Atoi(limitS)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code":    40002,
				"message": "Invalid request params!",
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
	currentPageS := ctx.Query("page")
	if currentPageS != "" {
		currentPage, err = strconv.Atoi(limitS)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code":    40002,
				"message": "Invalid request params!",
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
		err = models.Orm.Where("image_url <> '' AND video_url = ''").Offset(start).Limit(limit).Find(&jokes).Error
	case "video":
		err = models.Orm.Where("video_url <> '' AND image_url = ''").Offset(start).Limit(limit).Find(&jokes).Error
	case "text":
		err = models.Orm.Where("image_url = '' AND video_url = ''").Offset(start).Limit(limit).Find(&jokes).Error
	}

	if err != nil {
		logger.Error(err)
		ctx.JSON(500, gin.H{
			"code":    50001,
			"message": "connect error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    0,
		"data":    jokes,
		"message": "success",
	})
}

// 用于上拉刷新页面
func LoadRefresh(ctx *gin.Context) {

	var (
		err      error
		jokes    models.Jokes
		jokeType string
		limit    int
	)

	jokeType = ctx.Query("type")

	logger.Info("POST ==> 请求参数：", jokeType)

	// 判断type是否属于四种类型中的一种
	if _, ok := types[jokeType]; !ok {
		// 不存在
		ctx.JSON(200, gin.H{
			"code":    40002,
			"message": "Invalid request params!",
		})
		return
	}

	var payload PostLoadRefresh
	// 获取已在列表的 joke_id 集合、或者最新数据的时间
	ctx.BindJSON(&payload)

	if len(payload.JokeIds) <= 0 && payload.Time == "" {
		// 不存在
		ctx.JSON(200, gin.H{
			"code":    40002,
			"message": "Invalid request params!",
		})
		return
	}

	// 获取随机的数目

	rand.Seed(time.Now().UnixNano())
	limit = rand.Intn(15) + 1

	// 查询数据
	if payload.Time != "" {
		switch jokeType {
		case "hot":
			err = models.Orm.Where("time > ?", payload.Time).Limit(limit).Find(&jokes).Error
		case "pic":
			err = models.Orm.Where("time > ? AND image_url <> '' AND video_url = ''", payload.Time).Limit(limit).Find(&jokes).Error
		case "video":
			err = models.Orm.Where("time > ? AND video_url <> '' AND image_url = ''", payload.Time).Limit(limit).Find(&jokes).Error
		case "text":
			err = models.Orm.Where("time > ? AND image_url = '' AND video_url = ''", payload.Time).Limit(limit).Find(&jokes).Error
		}
	} else {
		switch jokeType {
		case "hot":
			err = models.Orm.Where(payload.JokeIds).Limit(limit).Find(&jokes).Error
		case "pic":
			err = models.Orm.Where(payload.JokeIds).Where("image_url <> '' AND video_url = ''").Limit(limit).Find(&jokes).Error
		case "video":
			err = models.Orm.Where(payload.JokeIds).Where("video_url <> '' OR image_url = ''").Limit(limit).Find(&jokes).Error
		case "text":
			err = models.Orm.Where(payload.JokeIds).Where("image_url = '' AND video_url = ''").Limit(limit).Find(&jokes).Error
		}
	}

	if err != nil {
		logger.Error(err)
		ctx.JSON(500, gin.H{
			"code":    50001,
			"message": "connect error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    0,
		"data":    jokes,
		"message": "success",
	})
}

// 获取糗事详情
func FindJokeInfo(ctx *gin.Context) {
	var (
		err  error
		joke models.Joke
		id   string
	)

	id = ctx.Param("id")

	logger.Info("糗事的ID ==> ", id)

	err = models.Orm.Where(&models.Joke{Id: id}).First(&joke).Error

	if err != nil {
		logger.Error(err)
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "id not exist!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    0,
		"data":    joke,
		"message": "success",
	})
}

// 获取糗事数量
func Count(ctx *gin.Context) {

	var count int

	err := models.Orm.Table("joke").Count(&count).Error

	logger.Info(count)
	if err != nil {
		// 不存在
		ctx.JSON(200, gin.H{
			"code":    40002,
			"message": "Invalid request params!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":  0,
		"count": count,
	})

}
