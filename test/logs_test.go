package test

import (
	"go.uber.org/zap"
	"joke-go/utils"
	"testing"
)

func TestLogZap(t *testing.T) {

	//logger, _ := zap.NewDevelopment()
	//
	//defer logger.Sync()
	//
	////logger.Info("adsda: ")
	//
	//sugar := logger.Sugar()
	//
	//mapD := map[string]interface{}{
	//	"demo": 12,
	//}
	//
	//sugar.Info("aaaa: ", mapD)

	//utils.Logger.Info("name: ")

	//utils.InitLogging()
	//utils.Log.Error("hahah", 12)

	utils.Logger.Info("aaaa: ", zap.String("name", "nana"))

	//fmt.Println(time.Hour*24)

}
