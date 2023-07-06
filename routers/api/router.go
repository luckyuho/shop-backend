package routers

import (
	"errors"
	"time"

	"basic/app/models/logInfo"
	v1 "basic/routers/api/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// query
	Gr1 := r.Group("/api/v1")
	Gr1.GET("/hello_world", v1.HelloWorld)
	Gr1.POST("/register", v1.ApiRegister)
	Gr1.POST("/login", v1.ApiLogin)
	Gr1.GET("/redirect", v1.ApiOauthCode2GetAccessToken)
	Gr1.GET("/products", v1.ApiGetAllProducts)
	// Gr1.POST("/purchase", v1.ApiPurchaseSql)
	Gr1.POST("/visa", v1.ApiPurchaseVisa)
	Gr1.POST("/notify", v1.ApiNotifyPurchase)
	Gr1.POST("/order", v1.ApiPushOrder)

	// 測試 mongodb 的使用情況
	logInfo.InsertLog("test", errors.New("this is a test"))
	// controllers.HandleOrder()  // 這裡有問題，要想一下要如何持續處理訂單，同時還要能起這個服務

	return r
}
