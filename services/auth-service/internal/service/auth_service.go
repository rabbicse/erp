package service

import (
	"errors"

	"github.com/rabbicse/auth-service/internal/domain"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (authService *AuthService) Authenticate(username string, password string) (*domain.User, error) {
	var user domain.User

	err := authService.db.
		Preload("Roles.Permissions").
		Where("username = ? AND is_active = true", username).
		First(&user).Error

	if err != nil {
		return nil, errors.New("Invalid credentials")
	}

	return &user, nil
}
