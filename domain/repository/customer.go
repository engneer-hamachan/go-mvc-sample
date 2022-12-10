package repository

import (
	"main/domain/model"
)

type CustomerRepository interface {
	GetCustomer(id int) (result *model.Customer, err error)
	GetCustomers() (result []model.Customer, err error)
	Create(c model.Customer) error
	Update(c model.Customer) error
	Delete(c model.Customer) error
}
