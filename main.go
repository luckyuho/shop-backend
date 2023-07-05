package main

import (
	"basic/app/models/logInfo"
	routers "basic/routers/api"
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	logInfo.InsertLog("test", errors.New("this is a test"))
	r := routers.InitRouter()
	r.Run(":9432")
}

func init() {
	// log 輸出為 json 格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 輸出設定為標準輸出(預設為 stderr)
	logrus.SetOutput(os.Stdout)
	// 設定要輸出的 log 等級
	logrus.SetLevel(logrus.DebugLevel)
}
