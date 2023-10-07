package services

import (
	"go-series-app/models"

	"gorm.io/gorm"
)

type Services interface {
	GetReport() models.Report
}

type services struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Services {
	return &services{
		db: db,
	}
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
