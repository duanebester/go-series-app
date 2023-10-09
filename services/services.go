package services

import (
	"go-series-app/models"

	"gorm.io/gorm"
)

type Services interface {
	GetReport() models.Report
	GetProduct() *models.Product
	GetUserByUsername(username string) (*models.User, error)
	UserService() UserService
}

type services struct {
	db          *gorm.DB
	userService UserService
}

func NewService(db *gorm.DB) Services {
	userService := NewUserService(db)
	return &services{
		db:          db,
		userService: userService,
	}
}

func (s *services) UserService() UserService {
	return s.userService
}

func (s *services) GetProduct() *models.Product {
	readProduct := &models.Product{}
	s.db.First(&readProduct, "code = ?", "D42")
	return readProduct
}

func (s *services) GetReport() models.Report {
	return models.Report{
		NumberUsers: 100,
	}
}

func (s *services) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := s.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
