package service

import (
	"context"
	"kstyle-test/entity"
	"kstyle-test/repository"
)

type LikeService struct {
	likeRepo   *repository.LikeRepository
	memberRepo *repository.MemberRepository
	reviewRepo *repository.ReviewRepository
}

func NewLikeService(likeRepo *repository.LikeRepository, memberRepo *repository.MemberRepository, reviewRepo *repository.ReviewRepository) *LikeService {
	return &LikeService{
		likeRepo:   likeRepo,
		memberRepo: memberRepo,
		reviewRepo: reviewRepo,
	}
}

func (m *LikeService) Like(ctx context.Context, like *entity.Like) (bool, error) {
	_, err := m.memberRepo.GetMemberByID(ctx, like.IDMember)
	if err != nil {
		return false, ErrMemberNotFound
	}
	_, err = m.reviewRepo.GetReviewByID(ctx, like.IDReview)
	if err != nil {
		return false, ErrReviewNotFound
	}
	res, err := m.likeRepo.Check(ctx, like)
	if err != nil {
		return false, err
	}
	if res {
		return false, m.likeRepo.DeleteLike(ctx, like)
	}
	return true, m.likeRepo.AddLike(ctx, like)
}
