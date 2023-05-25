package product

import (
	"basic/db"

	"github.com/sirupsen/logrus"
)

// 尋找資料
func FindProducts() []Product {

	var products []Product
	err := db.Get().Table("product").
		Find(&products).Error

	if err != nil {
		logrus.Error("讀取商品出錯", err)
	}

	return products
}

func UpdateProductStatus(
	id int,
	status bool,
) error {

	err := db.Get().Table("product").
		Model(&Product{}).
		Where(`id = ?`, id).
		Updates(map[string]interface{}{
			"status": status,
		}).Error

	return err
}
