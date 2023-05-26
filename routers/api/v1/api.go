package v1

import (
	"fmt"
	"net/http"

	ctr "basic/app/controllers"
	"basic/app/middlewares/auth"
	ProductModel "basic/app/models/product"
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
	result := ctr.RegisterUser(input.Name, input.Password)
	template(c, http.StatusOK, result)
}

// 使用者登入
func ApiLogin(c *gin.Context) {
	input := UserModel.User{}
	c.Bind(&input)
	result := ctr.LoginUser(input.Name, input.Password)
	template(c, http.StatusOK, result)
}

// 固定形式的輸出，有利於前端使用
func template(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"data": data,
	})
}

// 註冊使用者
func ApiOauthCode2GetAccessToken(c *gin.Context) {
	fullURL := c.Request.URL.String()
	status, accessToken := ctr.GetUserInfo(fullURL)

	if status == http.StatusBadRequest {
		template(c, status, accessToken)
	} else {
		token := ctr.GetCookie(c, accessToken.Id)
		if !token.Success {
			fmt.Println("錯誤")
			template(c, http.StatusBadRequest, "取得cookie時錯誤")
			return
		}
		// 指定目標網域的 URL
		targetURL := "http://localhost:8080/#/redirect?token=" + token.Token

		// 執行 URL 轉址到目標網域
		c.Redirect(http.StatusMovedPermanently, targetURL)
	}
}

func ApiGetAllProducts(c *gin.Context) {
	products := ctr.GetAllProducts()
	template(c, http.StatusOK, products)
}

// func ApiPurchaseSql(c *gin.Context) {
// 	input := ProductModel.Product{}
// 	c.Bind(&input)
// 	fmt.Println(input)
// 	err := ctr.PurchaseProduct(input.Id)
// 	if err != nil {
// 		template(c, http.StatusBadRequest, "購買商品時出錯")
// 	} else {
// 		template(c, http.StatusOK, nil)
// 	}
// }

func ApiPurchaseVisa(c *gin.Context) {
	input := ProductModel.GetId{}
	c.Bind(&input)
	fmt.Println(input)
	userId := auth.GetUserId(c)
	fmt.Println(userId)

	// data := ctr.PurchaseVisa(userId, input.Id)
	// template(c, http.StatusOK, userId)
}

type NewebPayReturn struct {
	Status     string
	MerchantID string
	Version    string
	TradeInfo  string
	TradeSha   string
}

func ApiNotifyPurchase(c *gin.Context) {
	input := NewebPayReturn{}
	c.Bind(&input)
	status, msg := ctr.ApiPurchaseResult(input.TradeInfo)

	template(c, status, msg)
}
