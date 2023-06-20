package entity

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	IDProduct  int    `json:"id_product"`
	IDMember   int    `json:"id_member"`
	DescReview string `json:"desc_review"`
	Likes      []Like `gorm:"foreignKey:IDReview"`
	Member     Member `gorm:"foreignKey:IDMember"`
}

type ReviewLikesCount struct {
	ReviewID   uint
	LikesCount int
}
