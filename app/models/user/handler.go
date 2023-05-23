package user

import (
	"basic/db"
	"fmt"
)

// 新增資料
func CreateUser(
	name,
	password string,
) error {
	user := User{Name: name, Password: password}
	err := db.Get().Table("user_table").
		Create(&user).Error
	return err
}

// 尋找資料
func FindUser(
	name string,
) error {
	user := User{Name: name}
	err := db.Get().Table("user_table").
		Where("name = ?", name).
		First(&user).Error
	fmt.Println(err)
	return err
}

// 從資料庫找有沒有相對應的 email 與 password
func LoginUser(name, password string) error {
	user := User{Name: name, Password: password}
	err := db.Get().Table("user_table").
		Where("name = ? and password = ?", name, password).
		Scan(&user).
		Error
	return err
}
