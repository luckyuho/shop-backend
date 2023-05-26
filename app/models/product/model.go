package product

type Product struct {
	Id        int    `gorm:"id" json:"id"`
	Name      string `gorm:"name" json:"name"`
	Price     int    `gorm:"price" json:"price"`
	Status    bool   `gorm:"status" json:"status"`
	Image     string `gorm:"image" json:"image"`
	TimeStamp int64  `gomr:"time_stamp" json:"time_stamp"`
}

type GetId struct {
	Id int `gorm:"id" json:"id"`
}

type Order struct {
	Id        int   `grom:"id" json:"id"`
	ProductId int   `gorm:"product_id" json:"product_id"`
	UserId    int   `grom:"user_id" json:"user_id"`
	TimeStamp int64 `gomr:"time_stamp" json:"time_stamp"`
}
