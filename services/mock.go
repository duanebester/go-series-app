package services

import "go-series-app/models"

type mock struct{}

func (m *mock) GetReport() models.Report {
	return models.Report{
		NumberUsers: 100,
	}
}

func NewMockService() Services {
	return &services{
		// set db conn
	}
}
