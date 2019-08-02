package utils

import (
	"bytes"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/gobuffalo/packr/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	EnvMode string
	Config  *ini.File
	err     error
)

func InitConfig() {

	// 参考：https://github.com/vicanso/articles/blob/master/viper.md?utm_source=tuicool&utm_medium=referral
	box := packr.New("config", "../config")

	configType := "yml"
	defaultConfig, err := box.Find("config.default.yml")

	if err != nil {
		Log.Panicln(err)
	}

	v := viper.New()
	v.SetConfigType(configType)
	err = v.ReadConfig(bytes.NewReader(defaultConfig))
	if err != nil {
		Log.Panicln("加载默认配置文件失败：\n", err)
		//return
	}
	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	// 获取程序启动参数
	envParam := flag.String("env", "dev", "--env dev | test | prod")
	flag.Parse()
	EnvMode = *envParam

	// 设置 gin 启动模式
	if EnvMode == "prod" {
		gin.SetMode(gin.ReleaseMode)
		Log.SetLevel(logrus.InfoLevel)
	} else if EnvMode == "test" {
		gin.SetMode(gin.TestMode)
		Log.SetLevel(logrus.InfoLevel)
	} else {
		gin.SetMode(gin.DebugMode)
		Log.SetLevel(logrus.InfoLevel)
	}

	Log.Info("Env：", EnvMode)

	var name = ""
	if EnvMode == "test" {
		name = "config.test.yml"
	} else if EnvMode == "prod" {
		name = "config.prod.yml"
	}

	// 根据配置的env读取相应的配置信息
	if name != "" {
		envConfig, err := box.Find(name)
		if err != nil {
			Log.Panicln(err)
		}

		viper.SetConfigType(configType)
		err = viper.ReadConfig(bytes.NewReader(envConfig))
		if err != nil {
			Log.Panicln("加载环境配置文件失败：\n", err)
		}
	}

}

// 获取配置文件参数
func GetConfig(key string) string {
	return viper.GetString(key)
}
