package service

import (
	"context"
	"kstyle-test/entity"
	"kstyle-test/repository"
)

type MemberService struct {
	memberRepo *repository.MemberRepository
}

func NewMemberService(memberRepo *repository.MemberRepository) *MemberService {
	return &MemberService{
		memberRepo: memberRepo,
	}
}

func (m *MemberService) GetAllMember(ctx context.Context) ([]entity.Member, error) {
	res, err := m.memberRepo.GetAllMember(ctx)
	if err != nil {
		return []entity.Member{}, err
	}
	return res, nil
}

func (m *MemberService) AddMember(ctx context.Context, member *entity.Member) error {
	err := m.memberRepo.AddMember(ctx, member)
	if err != nil {
		return err
	}
	return nil
}

func (m *MemberService) GetMemberByID(ctx context.Context, id int) (entity.Member, error) {
	res, err := m.memberRepo.GetMemberByID(ctx, id)
	if err != nil {
		return entity.Member{}, err
	}
	if res.ID == 0 {
		return entity.Member{}, ErrNotFound
	}

	return res, nil
}

func (m *MemberService) UpdateMember(ctx context.Context, member *entity.Member) (entity.Member, error) {
	err := m.memberRepo.UpdateMember(ctx, member)
	if err != nil {
		return entity.Member{}, err
	}

	return *member, nil
}

func (m *MemberService) DeleteMember(ctx context.Context, id int) error {
	err := m.memberRepo.DeleteMember(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
