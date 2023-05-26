package controllers

import (
	"basic/app/models/product"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/forgoer/openssl"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func GetAllProducts() []product.Product {
	return product.FindProducts()
}

func PurchaseProduct(id int) error {
	return product.UpdateProductStatus(id, false)
}

type VisaData struct {
	MerchantID  string
	TradeInfo   string
	TradeSha    string
	Version     string
	EncryptType int
}

type MpgInfo struct {
	merchantID string
	key        string
	iv         string
	notifyUrl  string
	returnUrl  string
	visaURL    string
	version    string
}

var mpgInfo MpgInfo

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("讀取環境變數錯誤", err)
	}

	mpgInfo.merchantID = os.Getenv("MerchantID")
	mpgInfo.key = os.Getenv("HashKey")
	mpgInfo.iv = os.Getenv("HashIv")
	mpgInfo.notifyUrl = os.Getenv("NotifyURL")
	mpgInfo.returnUrl = os.Getenv("ReturnURL")
	mpgInfo.visaURL = os.Getenv("VisaURL")
	mpgInfo.version = os.Getenv("Version")
}

func PurchaseVisa(userId, id int) VisaData {

	var visaData VisaData
	visaData.MerchantID = mpgInfo.merchantID
	visaData.TradeInfo = MpgRequest(id, userId)
	visaData.TradeSha = MpgSha(visaData.TradeInfo)
	visaData.Version = mpgInfo.version

	return visaData
}

func createOrderId() int64 {

	timeStamp := time.Now().Unix()

	return timeStamp
}

func createOrder(
	id,
	userId int,
) (product.Product, error) {

	info, err := product.FindProduct(id)
	if err != nil {
		return product.Product{}, err
	}

	timeStamp := createOrderId()
	info.TimeStamp = timeStamp

	_, err = product.CreateOrder(id, userId, timeStamp)
	if err != nil {
		return product.Product{}, err
	}

	return info, nil
}

func http_build_query(
	amount int,
	timeStamp int64,
	name string,
) string {

	reqData := "MerchantID=" + mpgInfo.merchantID
	reqData += "&RespondType=" + "JSON"
	reqData += "&TimeStamp=" + strconv.FormatInt(timeStamp, 10)
	reqData += "&Version=" + mpgInfo.version
	reqData += "&MerchantOrderNo=" + "Vanespl_ec_" + strconv.FormatInt(timeStamp, 10)
	reqData += "&Amt=" + strconv.Itoa(amount)
	reqData += "&NotifyURL=" + mpgInfo.notifyUrl
	reqData += "&ReturnURL=" + mpgInfo.returnUrl
	reqData += "&ItemDesc=" + name

	return reqData
}

func openssl_encrypt(reqData string) ([]byte, error) {

	req := []byte(reqData)
	key := []byte(mpgInfo.key)
	iv := []byte(mpgInfo.iv)
	dst, err := openssl.AesCBCEncrypt(req, key, iv, openssl.PKCS7_PADDING)
	if err != nil {
		logrus.Error("openssl_encrypt錯誤", err)
	}
	// fmt.Println(base64.StdEncoding.EncodeToString(dst)) // 1jdzWuniG6UMtoa3T6uNLA==

	// dst, _ = openssl.AesCBCDecrypt(dst, key, iv, openssl.PKCS7_PADDING)
	// fmt.Println("original: ", string(dst)) // 123456

	return dst, nil
}

func bin2hex(
	encrypt []byte,
) string {

	hexString := hex.EncodeToString(encrypt)
	return hexString

}

func MpgRequest(
	id int,
	userId int,
) string {

	info, err := createOrder(id, userId)
	if err != nil {
		return ""
	}

	reqData := http_build_query(info.Price, info.TimeStamp, info.Name)
	ciphertext, err := openssl_encrypt(reqData)
	if err != nil {
		logrus.Error("MpgRequest錯誤", err)
		return ""
	}

	hexString := bin2hex(ciphertext)
	return hexString

}

func MpgSha(
	tradeInfo string,
) string {

	hashs := "HashKey=" + mpgInfo.key + "&" + tradeInfo + "&HashIV=" + mpgInfo.iv

	hash256 := sha256.Sum256([]byte(hashs))

	hashString := hex.EncodeToString(hash256[:])
	upperHashString := strings.ToUpper(hashString)

	return upperHashString
}

type DecodeResult struct {
	Status string `json:"Status"`
	Result Result `json:"Result"`
}

type Result struct {
	MerchantID  string `json:"MerchantID"`
	TradeNo     string `json:"TradeNo"`
	PaymentType string `json:"PaymentType"`
}

func decode(hexString string) (string, error) {
	key := []byte(mpgInfo.key)
	iv := []byte(mpgInfo.iv)

	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		logrus.Error("解码失败:", err)
		return "", err
	}
	dst, _ := openssl.AesCBCDecrypt(decoded, key, iv, openssl.PKCS7_PADDING)

	var data map[string]interface{}
	err = json.Unmarshal(dst, &data)
	if err != nil {
		logrus.Error("解析 JSON 失败::", err)
		return "", err
	}

	if data["Status"].(string) == "SUCCESS" {
		result := data["Result"].(map[string]interface{})
		return result["MerchantOrderNo"].(string), nil
	}

	err = fmt.Errorf("訂單沒有成立")

	return "", err
}

func ApiPurchaseResult(hexString string) (int, error) {
	merchantOrderNo, err := decode(hexString)
	if err != nil {
		return http.StatusBadRequest, err
	}

	orderNo, err := strconv.ParseInt(merchantOrderNo[11:], 10, 64)
	if err != nil {
		logrus.Error("merchantOrderNo文字轉數字出錯", err)
		return http.StatusBadRequest, err
	}

	id, err := product.FindProductId(orderNo)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = PurchaseProduct(id)
	if err != nil {
		logrus.Error("更新購買商品時出錯", err)
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
