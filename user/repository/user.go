package repository

import (
	"context"
	"time"
	"user-services/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	columns []string
	db      *gorm.DB
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, req *ReqDeleteUser) error
	FindOne(ctx context.Context, req *ReqFirstUser, columns []string) (*model.User, error)
	Find(ctx context.Context, req *ReqFindAllUser, columns []string, skip, limit int) ([]model.User, error)
}

type ReqFirstUser struct {
	ID        uuid.UUID
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ReqDeleteUser struct {
	ID uuid.UUID
}

type ReqFindAllUser struct {
	IDs []uuid.UUID
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db:      db,
		columns: []string{"id", "email", "name", "password", "created_at", "updated_at", "deleted_at"},
	}
}

func (u *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := u.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Update(ctx context.Context, user *model.User) error {
	if err := u.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(ctx context.Context, req *ReqDeleteUser) error {
	if err := u.db.WithContext(ctx).Where(req).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) FindOne(ctx context.Context, req *ReqFirstUser, columns []string) (*model.User, error) {
	var user model.User
	query := u.db.WithContext(ctx)
	if len(columns) > 0 {
		query = query.Select(columns)
	} else {
		query = query.Select(u.columns)
	}
	if err := query.First(&user, req).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Find(ctx context.Context, req *ReqFindAllUser, columns []string, page, limit int) ([]model.User, error) {
	var users []model.User
	query := u.db.WithContext(ctx).Model(&model.User{})
	if len(columns) > 0 {
		query = query.Select(columns)
	} else {
		query = query.Select(u.columns)
	}
	if page > 0 {
		query = query.Offset((page - 1) * limit)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	if err := query.Find(&users, req).Error; err != nil {
		return nil, err
	}
	return users, nil
}
