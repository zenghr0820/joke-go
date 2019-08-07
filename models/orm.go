package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "joke-go/config"
	"joke-go/logger"
	"strconv"
)

var (
	Orm *gorm.DB
)

func InitOrm() {
	var err error

	// Mysql
	//Orm, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	GetConfig("database.user"),
	//	GetConfig("database.password"),
	//	GetConfig("database.host"),
	//	GetConfig("database.db")))

	// PostgreSQL
	Orm, err = gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		GetConfig("pg.host"),
		GetConfig("pg.user"),
		GetConfig("pg.db"),
		GetConfig("pg.password")))

	//Orm, err = pop.Connect(fmt.Sprintf("%s", GetConfig("database.env")))
	//defer Orm.Close()

	if err != nil {
		logger.Panic("err: ", err)
	}

	showSql, err := strconv.ParseBool(GetConfig("pg.showSql"))

	if err != nil {
		showSql = false
	}

	// 设置是否打印SQL
	Orm.LogMode(showSql)

	// 全局禁用表名复数
	Orm.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	// 自动创建/修改表结构
	//Orm.AutoMigrate(&Joke{})
	//Orm.AutoMigrate(&Author{})
	//Orm.AutoMigrate(&Book{})
}
