package services

import (
	"go-series-app/models"
	"go-series-app/repositories"
)

type ProductService interface {
	GetProductByID(id string) (*models.Product, error)
	CreateProduct(opts models.Product) (*models.Product, error)
	UpdateProduct(id string, opts models.Product) (*models.Product, error)
	DeleteProduct(id string) (*models.Product, error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (s *productService) GetProductByID(id string) (*models.Product, error) {
	return s.productRepository.GetProductByID(id)
}

func (s *productService) CreateProduct(opts models.Product) (*models.Product, error) {
	return s.productRepository.CreateProduct(opts)
}

func (s *productService) UpdateProduct(id string, opts models.Product) (*models.Product, error) {
	return s.productRepository.UpdateProduct(id, opts)
}

func (s *productService) DeleteProduct(id string) (*models.Product, error) {
	return s.productRepository.DeleteProduct(id)
}
