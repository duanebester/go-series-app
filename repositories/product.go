package repositories

import (
	"go-series-app/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(id string) (*models.Product, error)
	CreateProduct(opts models.Product) (*models.Product, error)
	UpdateProduct(id string, opts models.Product) (*models.Product, error)
	DeleteProduct(id string) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetProductByID(id string) (*models.Product, error) {
	product := &models.Product{}
	err := r.db.First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) CreateProduct(opts models.Product) (*models.Product, error) {
	product := models.NewProduct(&opts)
	err := r.db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) UpdateProduct(id string, updates models.Product) (*models.Product, error) {
	err := r.db.Model(&updates).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return nil, err
	}
	return r.GetProductByID(id)
}

func (r *productRepository) DeleteProduct(id string) (*models.Product, error) {
	product := &models.Product{}
	err := r.db.Delete(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
