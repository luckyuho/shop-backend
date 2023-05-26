package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func GetUserId(c *gin.Context) (userId int) {
	// 获取 Authorization 头部的值
	authHeader := c.GetHeader("Authorization")

	// 检查 Authorization 头部是否存在并且使用了 Bearer 认证方案
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		// 提取 JWT 的部分
		tokenString := authHeader[7:]

		// 解析 JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 在这里验证 JWT 的签名密钥（secret key）
			return []byte("test"), nil
		})

		if err == nil && token.Valid {
			// 提取 JWT 中的信息
			claims := token.Claims.(jwt.MapClaims)
			idConv, ok := claims["id"].(float64)
			id := int(idConv)
			if !ok {
				logrus.Error("Conversion failed")
				return 0
			}
			return id
		}

		return 0
	}
	return 0
}
