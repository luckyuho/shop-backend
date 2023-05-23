package v1

import (
	"net/http"

	UserCtr "basic/app/controllers/user"
	UserModel "basic/app/models/user"

	"github.com/gin-gonic/gin"
)

// 如果是要給前端的 api，函式只能有一個輸入參數 *gin.Context 且不能有輸出
func HelloWorld(c *gin.Context) {
	// 包成 json 的格式丟給前端
	c.JSON(http.StatusOK, gin.H{
		"data": "Hello world!",
	})
}

// 註冊使用者
func ApiRegister(c *gin.Context) {
	input := UserModel.User{}
	c.Bind(&input)
	result := UserCtr.RegisterUser(input.Email, input.Password)
	template(c, http.StatusOK, result)
}

// 使用者登入
func ApiLogin(c *gin.Context) {
	input := UserModel.User{}
	c.Bind(&input)
	result := UserCtr.LoginUser(input.Email, input.Password)
	template(c, http.StatusOK, result)
}

// 固定形式的輸出，有利於前端使用
func template(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"data": data,
	})
}
