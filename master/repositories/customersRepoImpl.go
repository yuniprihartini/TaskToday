package repositories

import (
	"database/sql"
	"fmt"
	"ujian/master/models"
	"ujian/tools"
)

// ServiceRepoImpl struct
type CustomersRepoImple struct {
	db *sql.DB
}

// GetAllDataCustomers method
func (s *CustomersRepoImple) GetAllDataCustomers() ([]*models.Customers, error) {
	var dataCustomers []*models.Customers
	data, err := s.db.Query(tools.GET_ALL_DATA_CUSTOMERS)
	if err != nil {
		return nil, err
	}

	for data.Next() {
		var Customers = models.Customers{}
		err := data.Scan(&Customers.CustomerId, &Customers.CustomerName, &Customers.BookingDate, &Customers.BookingTime,
			&Customers.NumberPlat, &Customers.ServiceID)
		if err != nil {
			return nil, err
		}
		dataCustomers = append(dataCustomers, &Customers)
	}
	return dataCustomers, nil
}

// GetDataById method
func (s *CustomersRepoImple) GetDataById(id string) (*models.Customers, error) {
	var dataCustomers = new(models.Customers)
	err := s.db.QueryRow(tools.GET_DATA_CUSTOMERS_BY_ID, id).Scan(&dataCustomers.CustomerId,
		&dataCustomers.CustomerName, &dataCustomers.BookingDate, &dataCustomers.BookingTime, &dataCustomers.NumberPlat,
		&dataCustomers.ServiceID)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return dataCustomers, nil
}

// UpdateDataCustomers method
func (s *CustomersRepoImple) UpdateDataCustomers(id string, Customers models.Customers) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.UPDATE_DATA_CUSTOMERS)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&Customers.CustomerName, &Customers.BookingDate, &Customers.BookingTime, &Customers.NumberPlat,
		&Customers.ServiceID, id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// InsertDataCustomers method
func (s *CustomersRepoImple) InsertDataCustomers(Customers models.Customers) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.INSERT_DATA_CUSTOMERS)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&Customers.CustomerId,
		&Customers.CustomerName, &Customers.BookingDate, &Customers.BookingTime, &Customers.NumberPlat,
		&Customers.ServiceID)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// DeleteCustomersData method
func (s *CustomersRepoImple) DeleteDataCustomersById(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(tools.DELETE_DATA_CUSTOMERS)
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
func InitCustomersRepoImpl(db *sql.DB) CustomersRepository {
	return &CustomersRepoImple{db}
}
