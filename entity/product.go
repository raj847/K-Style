package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	NameProduct string   `gorm:"type:varchar(255)" json:"name_product"`
	Price       float64  `json:"price"`
	Reviews     []Review `gorm:"foreignKey:IDProduct"`
}
