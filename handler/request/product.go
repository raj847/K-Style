package request

type Product struct {
	NameProduct string  `gorm:"type:varchar(255)" json:"name_product"`
	Price       float64 `json:"price"`
}
