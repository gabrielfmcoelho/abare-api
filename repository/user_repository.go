package repository

import (
	"context"

	"github.com/gabrielfmcoelho/abare-api/domain" // Update with your actual module path
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) error {
	result := ur.db.WithContext(ctx).Create(user)
	return result.Error
}

func (ur *userRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	result := ur.db.WithContext(ctx).
		Select("id", "name", "email", "created_at", "updated_at").
		Find(&users)
	return users, result.Error
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	result := ur.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user)
	return user, result.Error
}

func (ur *userRepository) GetByID(ctx context.Context, id uint) (domain.User, error) {
	var user domain.User
	result := ur.db.WithContext(ctx).
		First(&user, id)
	return user, result.Error
}
