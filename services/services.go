package services

import (
	"go-series-app/repositories"

	"gorm.io/gorm"
)

type Services interface {
	ReportService() ReportService
	ProductService() ProductService
	UserService() UserService
}

type services struct {
	db             *gorm.DB
	userService    UserService
	productService ProductService
	reportService  ReportService
}

func NewService(db *gorm.DB) Services {
	userService := NewUserService(db)
	reportService := NewReportService(db)
	productRepo := repositories.NewProductRepository(db)
	productService := NewProductService(productRepo)
	return &services{
		db:             db,
		userService:    userService,
		productService: productService,
		reportService:  reportService,
	}
}

func (s *services) UserService() UserService {
	return s.userService
}

func (s *services) ProductService() ProductService {
	return s.productService
}

func (s *services) ReportService() ReportService {
	return s.reportService
}
