package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Username  string   `db:"username"`
	Gender    string   `gorm:"type:varchar(255)" db:"gender"`
	Skintype  string   `gorm:"type:varchar(255)" db:"skintype"`
	Skincolor string   `gorm:"type:varchar(255)" db:"skincolor"`
	Reviews   []Review `gorm:"foreignKey:IDMember"`
	Likes     []Like   `gorm:"foreignKey:IDMember"`
}
