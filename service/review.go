package service

import (
	"context"
	"errors"
	"kstyle-test/entity"
	"kstyle-test/repository"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrMemberNotFound  = errors.New("member not found")
	ErrProductNotFound = errors.New("product not found")
	ErrReviewNotFound  = errors.New("review not found")
)

type ReviewService struct {
	reviewRepo  *repository.ReviewRepository
	productRepo *repository.ProductRepository
	memberRepo  *repository.MemberRepository
}

func NewReviewService(reviewRepo *repository.ReviewRepository, memberRepo *repository.MemberRepository, productRepo *repository.ProductRepository) *ReviewService {
	return &ReviewService{
		reviewRepo:  reviewRepo,
		memberRepo:  memberRepo,
		productRepo: productRepo,
	}
}

func (m *ReviewService) AddReview(ctx context.Context, review *entity.Review) error {
	_, err := m.memberRepo.GetMemberByID(ctx, review.IDMember)
	if err != nil {
		return ErrMemberNotFound
	}
	_, err = m.productRepo.GetProductByID(ctx, review.IDProduct)
	if err != nil {
		return ErrProductNotFound
	}
	err = m.reviewRepo.AddReview(ctx, review)
	if err != nil {
		return err
	}
	return nil
}
