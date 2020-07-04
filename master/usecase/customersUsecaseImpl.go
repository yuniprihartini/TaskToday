package usecase

import (
	"ujian/master/models"
	"ujian/master/repositories"
	"ujian/tools"
)

type CustomersUsecaseImpl struct {
	CustomersRepo repositories.CustomersRepository
}

func (s CustomersUsecaseImpl) GetAllDataCustomers() ([]*models.Customers, error) {
	Customers, err := s.CustomersRepo.GetAllDataCustomers()
	if err != nil {
		return nil, err
	}

	return Customers, nil
}

func (s CustomersUsecaseImpl) GetDataById(id string) (*models.Customers, error) {
	Customers, err := s.CustomersRepo.GetDataById(id)
	if err != nil {
		return nil, err
	}

	return Customers, nil
}
func (s CustomersUsecaseImpl) InsertDataCustomers(Customers models.Customers) error {
	err := s.CustomersRepo.InsertDataCustomers(Customers)
	tools.FatalErr(err)

	return nil
}
func (s CustomersUsecaseImpl) UpdateDataCustomers(id string, Customers models.Customers) error {
	err := s.CustomersRepo.UpdateDataCustomers(id, Customers)
	tools.FatalErr(err)

	return nil
}
func (s CustomersUsecaseImpl) DeleteDataCustomersById(id string) error {
	err := s.CustomersRepo.DeleteDataCustomersById(id)
	tools.FatalErr(err)

	return nil
}
func InitCustomersUseCaseImpl(CustomersRepo repositories.CustomersRepository) CustomersUseCase {
	return &CustomersUsecaseImpl{CustomersRepo}
}
