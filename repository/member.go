package repository

import (
	"context"
	"kstyle-test/entity"

	"gorm.io/gorm"
)

type MemberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{db}
}

func (m *MemberRepository) GetAllMember(ctx context.Context) ([]entity.Member, error) {
	var members []entity.Member

	err := m.db.
		WithContext(ctx).
		Find(&members).Error
	if err != nil {
		return []entity.Member{}, err
	}

	return members, nil
}

func (m *MemberRepository) AddMember(ctx context.Context, member *entity.Member) error {
	err := m.db.
		WithContext(ctx).
		Create(&member).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *MemberRepository) GetMemberByID(ctx context.Context, id int) (entity.Member, error) {
	var member entity.Member

	err := m.db.
		WithContext(ctx).
		Table("members").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&member).Error
	if err != nil {
		return entity.Member{}, err
	}

	return member, nil
}

func (m *MemberRepository) DeleteMember(ctx context.Context, id int) error {
	err := m.db.
		WithContext(ctx).
		Delete(&entity.Member{}, id).Error
	return err
}

func (m *MemberRepository) UpdateMember(ctx context.Context, member *entity.Member) error {
	err := m.db.
		WithContext(ctx).
		Table("members").
		Where("id = ?", member.ID).
		Updates(&member).Error
	return err
}
