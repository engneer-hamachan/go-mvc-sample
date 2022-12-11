package repository

import (
	"main/domain/model/customer"
)

type CustomerRepository interface {
	GetCustomerByEmail(email string) (result *customer.Customer, err error)
	GetCustomer(id string) (result *customer.Customer, err error)
	GetCustomers() (result []*customer.Customer, err error)
	InsertCustomer(c *customer.Customer) error
	UpdateCustomer(c *customer.Customer) error
	DeleteCustomer(c *customer.Customer) error
}
