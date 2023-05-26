package user

type User struct {
	Id       int    `gorm:"id" json:"id"`
	Name     string `gorm:"name" json:"name"`
	Password string `gorm:"password" json:"password"`
}
