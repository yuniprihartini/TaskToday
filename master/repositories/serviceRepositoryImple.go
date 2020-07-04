package repositories

import (
	"database/sql"
	"fmt"
	"ujian/master/models"
	"ujian/tools"
)

// ServiceRepoImpl struct
type ServicesRepoImple struct {
	db *sql.DB
}

// GetAllDataService method
func (s *ServicesRepoImple) GetAllDataService() ([]*models.Services, error) {
	var dataServices []*models.Services
	data, err := s.db.Query(tools.GET_ALL_DATA_SERVICES)
	if err != nil {
		return nil, err
	}

	for data.Next() {
		var Services = models.Services{}
		err := data.Scan(&Services.ServiceId, &Services.ServiceType, &Services.VehicleType, &Services.ServicePrice)
		if err != nil {
			return nil, err
		}
		dataServices = append(dataServices, &Services)
	}
	return dataServices, nil
}

// GetDataById method
func (s *ServicesRepoImple) GetDataById(id string) (*models.Services, error) {
	var dataServices = new(models.Services)
	err := s.db.QueryRow(tools.GET_DATA_SERVICES_BY_ID, id).Scan(&dataServices.ServiceId,
		&dataServices.ServiceType, &dataServices.VehicleType, &dataServices.ServicePrice)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return dataServices, nil
}

// UpdateDataService method
func (s *ServicesRepoImple) UpdateDataService(id string, Services models.Services) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.UPDATE_DATA_SERVICES)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&Services.ServiceType, &Services.VehicleType, &Services.ServicePrice, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// InsertDataService method
func (s *ServicesRepoImple) InsertDataService(Services models.Services) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.INSERT_DATA_SERVICES)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&Services.ServiceId, &Services.ServiceType, &Services.VehicleType, &Services.ServicePrice)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// DeleteServiceData method
func (s *ServicesRepoImple) DeleteDataServiceById(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.DELETE_DATA_SERVICES)
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
func InitServiceRepoImpl(db *sql.DB) ServicesRepository {
	return &ServicesRepoImple{db}
}
