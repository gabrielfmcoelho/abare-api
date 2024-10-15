package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	// Gorm model includes ID, CreatedAt, UpdatedAt, DeletedAt fields
	gorm.Model
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;uniqueIndex;not null"`
	Password string `gorm:"size:255;not null"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Fetch(ctx context.Context) ([]User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id uint) (User, error)
}
