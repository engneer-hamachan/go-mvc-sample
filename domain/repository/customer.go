package repository

import (
	"main/domain/model/customer"
)

type CustomerRepository interface {
	GetCustomer(id string) (result *customer.Customer, err error)
	GetCustomers() (result []customer.Customer, err error)
	Create(c *customer.Customer) error
	Update(c *customer.Customer) error
	Delete(c *customer.Customer) error
}
