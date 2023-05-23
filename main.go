package main

import (
	routers "basic/routers/api"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
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
