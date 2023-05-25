package controllers

import (
	"basic/app/models/product"
	"os"

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

func PurchaseVisa(id int) VisaData {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error(err)
	}

	var visaData VisaData
	visaData.MerchantID = os.Getenv("MerchantID")
	visaData.TradeInfo = "111"
	visaData.TradeSha = "222"
	visaData.Version = "333"
	visaData.EncryptType = 1

	return visaData
}
