package repositories

import "ujian/master/models"

type ProgressRepository interface {
	GetAllDataProgress() ([]*models.Progress, error)
	GetDataById(id string) (*models.Progress, error)
	InsertDataProgress(service models.Progress) error
	UpdateDataProgress(id string, service models.Progress) error
	DeleteDataProgressById(id string) error
}
