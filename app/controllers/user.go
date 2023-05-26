package controllers

import (
	UserModel "basic/app/models/user"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

func CreateJwtToken(
	id int,
) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 過期時間為 24 小時後
	}

	// 使用密鑰簽署 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWTSECRET")

	// 使用密鑰生成最終的 JWT 字串
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

// 註冊使用者
// 這邊可以重複註冊，如果不希望有重複的 email 出現，可在資料庫中設定 email 屬性為 unique
type token struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

func RegisterUser(
	name,
	password string,
) token {

	user, err := UserModel.CreateUser(name, password)
	if err != nil {
		return token{
			Success: false,
			Token:   "",
		}
	}

	tokenString, err := CreateJwtToken(user.Id)

	return token{
		Success: err == nil,
		Token:   tokenString,
	}
}

// 使用者登入
// 如果資料庫中沒有找到對應的使用者帳密，回傳 err = record not found，有找到則 err = nil
func LoginUser(
	name,
	password string,
) token {
	user, err := UserModel.LoginUser(name, password)

	if err != nil {
		return token{
			Success: false,
			Token:   "",
		}
	}

	tokenString, err := CreateJwtToken(user.Id)

	return token{
		Success: err == nil,
		Token:   tokenString,
	}
}
