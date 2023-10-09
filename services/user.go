package services

import (
	"go-series-app/models"

	"gorm.io/gorm"
)

type UserService interface {
	GetUserByUsername(username string) (*models.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (s *userService) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := s.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
