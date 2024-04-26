package repository

import (
	"github.com/9-Spontan-ZooPort/9_Spontan_ZooPort-BE/internal/pkg/entity"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (entity.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *AuthRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
