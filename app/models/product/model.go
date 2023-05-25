package product

type Product struct {
	Id     int    `gorm:"id" json:"id"`
	Name   string `gorm:"name" json:"name"`
	Price  int    `gorm:"price" json:"price"`
	Status bool   `gorm:"status" json:"status"`
	Image  string `gorm:"image" json:"image"`
}
