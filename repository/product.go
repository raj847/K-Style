package repository

import (
	"context"
	"kstyle-test/entity"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (m *ProductRepository) GetAllProduct(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product

	err := m.db.Model(&entity.Product{}).
		Preload("Reviews").
		Preload("Reviews.Member").
		Find(&products).Error
	if err != nil {
		return []entity.Product{}, err
	}
	return products, nil
}

func (m *ProductRepository) GetAllLikes(ctx context.Context) ([]entity.ReviewLikesCount, error) {
	var reviewLikesCount []entity.ReviewLikesCount
	err := m.db.Table("reviews").
		Select("reviews.id AS review_id, COUNT(likes.id) AS likes_count").
		Joins("LEFT JOIN likes ON likes.id_review = reviews.id").
		Group("reviews.id").
		Scan(&reviewLikesCount).Error
	if err != nil {
		return []entity.ReviewLikesCount{}, err
	}
	return reviewLikesCount, nil
}

func (m *ProductRepository) AddProduct(ctx context.Context, product *entity.Product) error {
	err := m.db.
		WithContext(ctx).
		Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *ProductRepository) GetProductByID(ctx context.Context, id int) (entity.Product, error) {
	var product entity.Product

	err := m.db.Model(&entity.Product{}).
		Preload("Reviews").
		Preload("Reviews.Member").
		Where("id=?", id).
		First(&product).Error
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (m *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	err := m.db.
		WithContext(ctx).
		Delete(&entity.Product{}, id).Error
	return err
}

func (m *ProductRepository) UpdateProduct(ctx context.Context, product *entity.Product) error {
	err := m.db.
		WithContext(ctx).
		Table("products").
		Where("id = ?", product.ID).
		Updates(&product).Error
	return err
}
