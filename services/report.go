package services

import (
	"go-series-app/models"

	"gorm.io/gorm"
)

type ReportService interface {
	GetReport() models.Report
}

type reportService struct {
	db *gorm.DB
}

func NewReportService(db *gorm.DB) ReportService {
	return &reportService{
		db: db,
	}
}

func (s *reportService) GetReport() models.Report {
	return models.Report{
		NumberUsers: 100,
	}
}
