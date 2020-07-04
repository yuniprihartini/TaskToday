package repositories

import "ujian/master/models"

type CustomersRepository interface {
	GetAllDataCustomers() ([]*models.Customers, error)
	GetDataById(id string) (*models.Customers, error)
	InsertDataCustomers(service models.Customers) error
	UpdateDataCustomers(id string, service models.Customers) error
	DeleteDataCustomersById(id string) error
}
