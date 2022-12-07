package repository

import (
	"main/domain/model"
)

type CustomerRepository interface {
	GetCustomer(id int) model.Customer
	GetCustomers() []model.Customer
	Create(c model.Customer)
	Update(c model.Customer)
	Delete(c model.Customer)
}
