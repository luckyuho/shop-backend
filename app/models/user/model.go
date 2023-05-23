package user

type User struct {
	Email    string `gorm:"email" json:"email"`
	Password string `gorm:"password" json:"password"`
}
