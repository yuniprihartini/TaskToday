package usecase

import "ujian/master/models"

type ServicesUseCase interface {
	GetAllDataService() ([]*models.Services, error)
	GetDataById(id string) (*models.Services, error)
	InsertDataService(service models.Services) error
	UpdateDataService(id string, service models.Services) error
	DeleteDataServiceById(id string) error
}
