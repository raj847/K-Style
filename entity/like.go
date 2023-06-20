package entity

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	IDReview int `json:"id_review" gorm:"uniqueIndex:idx_member_review"`
	IDMember int `json:"id_member" gorm:"uniqueIndex:idx_member_review"`
}
