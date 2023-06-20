package repository

import (
	"context"
	"kstyle-test/entity"

	"gorm.io/gorm"
)

type LikeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{db}
}

func (m *LikeRepository) AddLike(ctx context.Context, like *entity.Like) error {
	err := m.db.
		WithContext(ctx).
		Create(&like).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *LikeRepository) DeleteLike(ctx context.Context, like *entity.Like) error {
	err := m.db.
		WithContext(ctx).
		Unscoped().
		Delete(&entity.Like{}, like).Error
	return err
}

func (m *LikeRepository) Check(ctx context.Context, like *entity.Like) (bool, error) {
	result := m.db.WithContext(ctx).First(&entity.Like{}, like)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}
