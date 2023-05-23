package routers

import (
	"time"

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
	return r
}
