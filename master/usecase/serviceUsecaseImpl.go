package usecase

import (
	"ujian/master/models"
	"ujian/master/repositories"
	"ujian/tools"
)

type ServicesUsecaseImpl struct {
	ServicesRepo repositories.ServicesRepository
}

func (s ServicesUsecaseImpl) GetAllDataService() ([]*models.Services, error) {
	Services, err := s.ServicesRepo.GetAllDataService()
	if err != nil {
		return nil, err
	}

	return Services, nil
}

func (s ServicesUsecaseImpl) GetDataById(id string) (*models.Services, error) {
	Services, err := s.ServicesRepo.GetDataById(id)
	if err != nil {
		return nil, err
	}

	return Services, nil
}
func (s ServicesUsecaseImpl) InsertDataService(Services models.Services) error {
	err := s.ServicesRepo.InsertDataService(Services)
	tools.FatalErr(err)

	return nil
}
func (s ServicesUsecaseImpl) UpdateDataService(id string, Services models.Services) error {
	err := s.ServicesRepo.UpdateDataService(id, Services)
	tools.FatalErr(err)

	return nil
}
func (s ServicesUsecaseImpl) DeleteDataServiceById(id string) error {
	err := s.ServicesRepo.DeleteDataServiceById(id)
	tools.FatalErr(err)

	return nil
}
func InitServiceUseCaseImpl(ServicesRepo repositories.ServicesRepository) ServicesUseCase {
	return &ServicesUsecaseImpl{ServicesRepo}
}
