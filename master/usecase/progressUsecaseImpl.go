package usecase

import (
	"ujian/master/models"
	"ujian/master/repositories"
	"ujian/tools"
)

type ProgressUsecaseImpl struct {
	ProgressRepo repositories.ProgressRepository
}

func (s ProgressUsecaseImpl) GetAllDataProgress() ([]*models.Progress, error) {
	Progress, err := s.ProgressRepo.GetAllDataProgress()
	if err != nil {
		return nil, err
	}

	return Progress, nil
}

func (s ProgressUsecaseImpl) GetDataById(id string) (*models.Progress, error) {
	Progress, err := s.ProgressRepo.GetDataById(id)
	if err != nil {
		return nil, err
	}

	return Progress, nil
}
func (s ProgressUsecaseImpl) InsertDataProgress(Progress models.Progress) error {
	err := s.ProgressRepo.InsertDataProgress(Progress)
	tools.FatalErr(err)

	return nil
}
func (s ProgressUsecaseImpl) UpdateDataProgress(id string, Progress models.Progress) error {
	err := s.ProgressRepo.UpdateDataProgress(id, Progress)
	tools.FatalErr(err)

	return nil
}
func (s ProgressUsecaseImpl) DeleteDataProgressById(id string) error {
	err := s.ProgressRepo.DeleteDataProgressById(id)
	tools.FatalErr(err)

	return nil
}
func InitProgressUseCaseImpl(ProgressRepo repositories.ProgressRepository) ProgressUseCase {
	return &ProgressUsecaseImpl{ProgressRepo}
}
