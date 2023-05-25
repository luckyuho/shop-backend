package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	UserModel "basic/app/models/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func GetUserInfo(rawURL string) (int, UserInfo) {
	code := GetOauthCode(rawURL)
	status, accessToken := CodeToAccessToken(code)

	if status == http.StatusBadRequest {
		return status, accessToken
	}

	status, token := GetUser(accessToken.AccessToken)

	if UserModel.FindUser(token.IdString) != nil {
		UserModel.CreateUser(token.IdString, "github")
	}

	return status, token
}

func GetOauthCode(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "解析 URL 錯誤"
	}

	// 獲取 URL 中的查詢參數
	queryParams := parsedURL.Query()

	// 獲取特定參數的值
	code := queryParams.Get("code")

	return code
}

type ResponseMsg struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func CodeToAccessToken(code string) (int, UserInfo) {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error(err)
		return http.StatusOK, UserInfo{ErrorMsg: fmt.Sprintf("讀取 .env 發生問題: %s", err)}
	}

	accessUrl := "https://github.com/login/oauth/access_token"

	formValues := url.Values{}
	formValues.Set("grant_type", "authorization_code")
	formValues.Set("client_id", os.Getenv("GITCLIENTID"))
	formValues.Set("client_secret", os.Getenv("GITSECRET"))
	formValues.Set("code", code)
	formValues.Set("redirect_uri", os.Getenv("REDIRECTURL"))
	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)

	req, err := http.NewRequest("POST", accessUrl, formBytesReader)
	if err != nil {
		return http.StatusBadRequest, UserInfo{ErrorMsg: fmt.Sprintf("創建 POST 請求錯誤: %s", err)}
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	//Do方法发送请求
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return http.StatusBadRequest, UserInfo{ErrorMsg: fmt.Sprintf("oauth2 拿 code 取 token 錯誤: %s", err)}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return http.StatusBadRequest, UserInfo{ErrorMsg: fmt.Sprintf("讀取 Api 回來的 token 時發生錯誤: %s", err)}
	}
	// fmt.Println(string(body))

	var result ResponseMsg
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		return http.StatusBadRequest, UserInfo{ErrorMsg: fmt.Sprintf("無法解析 access token 的 JSON: %s", err)}
	}

	// logrus.Info("user_token:", result.AccessToken)
	return http.StatusOK, UserInfo{AccessToken: result.AccessToken}
}

type UserInfo struct {
	Login       string `json:"login"`
	Id          int    `json:"id"`
	IdString    string `json:"id_string"`
	ErrorMsg    string `json:"error_msg"`
	AccessToken string `json:"access_token"`
}

func GetUser(accessToken string) (int, UserInfo) {
	getUserUrl := "https://api.github.com/user"

	req, err := http.NewRequest("GET", getUserUrl, nil)
	if err != nil {
		return http.StatusBadRequest, UserInfo{ErrorMsg: fmt.Sprintf("創建 GET 請求錯誤: %s", err)}
	}

	BearerToken := "Bearer " + accessToken
	req.Header.Set("Authorization", BearerToken)

	//Do方法发送请求
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return http.StatusBadRequest, UserInfo{ErrorMsg: fmt.Sprintf("access_token 拿 使用者資訊 錯誤: %s", err)}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return http.StatusBadRequest, UserInfo{ErrorMsg: fmt.Sprintf("讀取 Api 回來的 使用者資訊 時發生錯誤: %s", err)}
	}
	// fmt.Println(string(body))

	var result UserInfo
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		return http.StatusBadRequest, UserInfo{ErrorMsg: fmt.Sprintf("無法解析 使用者資訊 的 JSON: %s", err)}
	}

	return http.StatusOK, UserInfo{Login: result.Login, Id: result.Id, IdString: strconv.Itoa(result.Id)}
}

func GetCookie(c *gin.Context, name string) token {
	tokenString, err := CreateJwtToken(name)

	return token{
		Success: err == nil,
		Token:   tokenString,
	}
}
