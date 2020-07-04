package repositories

import (
	"database/sql"
	"fmt"
	"ujian/master/models"
	"ujian/tools"
)

// ServiceRepoImpl struct
type ProgressRepoImple struct {
	db *sql.DB
}

// GetAllDataProgress method
func (s *ProgressRepoImple) GetAllDataProgress() ([]*models.Progress, error) {
	var dataProgress []*models.Progress
	data, err := s.db.Query(tools.GET_ALL_PROGRESS)
	if err != nil {
		return nil, err
	}

	for data.Next() {
		var Progress = models.Progress{}
		err := data.Scan(&Progress.ProgressId, &Progress.ServiceId, &Progress.CustomerId, &Progress.ServiceStatus)
		if err != nil {
			return nil, err
		}
		dataProgress = append(dataProgress, &Progress)
	}
	return dataProgress, nil
}

// GetDataById method
func (s *ProgressRepoImple) GetDataById(id string) (*models.Progress, error) {
	var dataProgress = new(models.Progress)
	err := s.db.QueryRow(tools.GET_DATA_PROGRESS_BY_ID, id).Scan(&dataProgress.ProgressId, &dataProgress.ServiceId, &dataProgress.CustomerId,
		&dataProgress.ServiceStatus)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return dataProgress, nil
}

// UpdateDataProgress method
func (s *ProgressRepoImple) UpdateDataProgress(id string, Progress models.Progress) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.UPDATE_DATA_PROGRESS)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&Progress.ServiceId, &Progress.CustomerId, &Progress.ServiceStatus, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// InsertDataProgress method
func (s *ProgressRepoImple) InsertDataProgress(Progress models.Progress) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.INSERT_DATA_PROGRESS)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&Progress.ProgressId, &Progress.ServiceId, &Progress.CustomerId, &Progress.ServiceStatus)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// DeleteProgressData method
func (s *ProgressRepoImple) DeleteDataProgressById(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.DELETE_DATA_PROGRESS)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// TemporaryEmployeeRepoImpl init
func InitProgressRepoImpl(db *sql.DB) ProgressRepository {
	return &ProgressRepoImple{db}
}
