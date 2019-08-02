package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"joke-go/models"
	"strconv"
)

var (
	Orm *gorm.DB
)

func InitOrm() {
	var err error

	Orm, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetConfig("database.user"),
		GetConfig("database.password"),
		GetConfig("database.host"),
		GetConfig("database.db")))

	//Orm, err = pop.Connect(fmt.Sprintf("%s", GetConfig("database.env")))
	//defer Orm.Close()

	if err != nil {
		Log.Panicln(err)
	}

	showSql, err := strconv.ParseBool(GetConfig("database.showSql"))

	if err != nil {
		showSql = false
	}

	// 设置是否打印SQL
	Orm.LogMode(showSql)

	// 全局禁用表名复数
	Orm.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	// 自动创建/修改表结构
	Orm.AutoMigrate(&models.Joke{})
	Orm.AutoMigrate(&models.Author{})
	Orm.AutoMigrate(&models.Book{})
}
