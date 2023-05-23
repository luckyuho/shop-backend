package user

import (
	"basic/db"
)

// 新增資料
func CreateUser(
	email,
	password string,
) error {
	user := User{Email: email, Password: password}
	err := db.Get().Table("user_table").
		Create(&user).Error
	return err
}

// 從資料庫找有沒有相對應的 email 與 password
func LoginUser(email, password string) error {
	user := User{Email: email, Password: password}
	err := db.Get().Table("user_table").
		Where("email = ? and password = ?", email, password).
		Scan(&user).
		Error
	return err
}
