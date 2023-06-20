package repository

import (
	"context"
	"kstyle-test/entity"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db}
}

func (m *ReviewRepository) AddReview(ctx context.Context, review *entity.Review) error {
	err := m.db.
		WithContext(ctx).
		Create(&review).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *ReviewRepository) GetReviewByID(ctx context.Context, id int) (entity.Review, error) {
	var review entity.Review

	err := m.db.
		WithContext(ctx).
		Table("reviews").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&review).Error
	if err != nil {
		return entity.Review{}, err
	}

	return review, nil
}
