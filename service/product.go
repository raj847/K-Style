package service

import (
	"context"
	"kstyle-test/entity"
	"kstyle-test/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (m *ProductService) GetAllProduct(ctx context.Context) ([]entity.Product, error) {
	res, err := m.productRepo.GetAllProduct(ctx)
	if err != nil {
		return []entity.Product{}, err
	}
	return res, nil
}

func (m *ProductService) AddProduct(ctx context.Context, product *entity.Product) error {
	err := m.productRepo.AddProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (m *ProductService) GetProductByID(ctx context.Context, id int) (entity.Product, error) {
	res, err := m.productRepo.GetProductByID(ctx, id)
	if err != nil {
		return entity.Product{}, err
	}
	return res, nil
}

func (m *ProductService) UpdateProduct(ctx context.Context, product *entity.Product) (entity.Product, error) {
	err := m.productRepo.UpdateProduct(ctx, product)
	if err != nil {
		return entity.Product{}, err
	}

	return *product, nil
}

func (m *ProductService) DeleteProduct(ctx context.Context, id int) error {
	err := m.productRepo.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *ProductService) GetAllLikes(ctx context.Context) ([]entity.ReviewLikesCount, error) {
	return m.productRepo.GetAllLikes(ctx)
}
