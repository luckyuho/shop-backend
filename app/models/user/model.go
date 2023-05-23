package user

type User struct {
	Name     string `gorm:"name" json:"name"`
	Password string `gorm:"password" json:"password"`
}
