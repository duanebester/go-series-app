package services

import (
	"go-series-app/models"

	"gorm.io/gorm"
)

type ProductService interface {
	GetProductByID(id string) (*models.Product, error)
}

type productService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{
		db: db,
	}
}

func (s *productService) GetProductByID(id string) (*models.Product, error) {
	product := &models.Product{}
	err := s.db.Find(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
