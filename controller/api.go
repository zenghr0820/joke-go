package controller

import (
	"github.com/gin-gonic/gin"
	"joke-go/models"
	. "joke-go/utils"
	"time"
)

var types = map[string]bool{
	"text":  true,
	"pic":   true,
	"video": true,
	"hot":   true,
}

func Fetch(ctx *gin.Context) {

	var (
		err error
		jokes models.Jokes
	)

	jokeType := ctx.Param("type")
	Log.Println(" --- 请求参数：", jokeType)
	if types[jokeType] {
		err = Orm.Where("type = ?", jokeType).Find(&jokes).Error
		if err != nil {
			Log.Errorln(err)
			ctx.JSON(500, gin.H{
				"code": 50001,
				"msg":  "connect error",
			})
			return
		}
	} else {
		ctx.JSON(500, gin.H{
			"code": 40002,
			"msg":  "type error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": jokes,
		"msg":  "success",
	})
}

func Create(ctx *gin.Context)  {

	var (
		err error
		joke models.Joke
	)


	joke.Id = "2018"
	joke.Type = "hot"
	joke.Title = "cehsi"
	joke.Time = time.Now()

	err = Orm.Create(&joke).Error

	if err != nil {
		Log.Errorln(err)
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
