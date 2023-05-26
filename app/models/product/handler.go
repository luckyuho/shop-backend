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

func FindProduct(id int) (Product, error) {
	var product Product
	err := db.Get().Table("product").
		Where(`id = ?`, id).
		First(&product).Error

	if err != nil {
		logrus.Error("讀取商品出錯", err)
		return Product{}, err
	}

	return product, nil
}

func CreateOrder(
	id,
	userId int,
	timeStamp int64,
) (Order, error) {

	order := Order{ProductId: id, UserId: userId, TimeStamp: timeStamp}
	err := db.Get().Table("order").
		Create(&order).Error

	return order, err
}

func FindProductId(
	TimeStamp int64,
) (int, error) {

	var order Order
	err := db.Get().Table("order").
		Where(`time_stamp = ?`, TimeStamp).
		First(&order).Error

	if err != nil {
		logrus.Error("讀取訂單出錯", err)
		return 0, err
	}
	return order.ProductId, nil
}
