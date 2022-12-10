package usecase

import (
	"main/domain/model"
	"main/domain/repository"
)

type CustomerUseCase interface {
	GetCustomer(id int) (result *model.Customer, err error)
	GetCustomers() (result []model.Customer, err error)
	CreateCustomer(name string, age int) error
	UpdateCustomer(id int, name string, age int) error
	DeleteCustomer(id int) error
}

type customerUseCase struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerUseCase(cr repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepository: cr,
	}
}

func (cu *customerUseCase) GetCustomer(id int) (result *model.Customer, err error) {
	customer, err := cu.customerRepository.GetCustomer(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (cu *customerUseCase) GetCustomers() (result []model.Customer, err error) {
	customers, err := cu.customerRepository.GetCustomers()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (cu *customerUseCase) CreateCustomer(name string, age int) error {
	customer := model.Customer{Name: name, Age: age}
	err := cu.customerRepository.Create(customer)
	if err != nil {
		return err
	}

	return nil
}

func (cu *customerUseCase) UpdateCustomer(id int, name string, age int) error {
	customer, err := cu.customerRepository.GetCustomer(id)
	if err != nil {
		return err
	}

	customer.Name = name
	customer.Age = age
	err = cu.customerRepository.Update(*customer)
	if err != nil {
		return err
	}

	return nil
}

func (cu *customerUseCase) DeleteCustomer(id int) error {
	customer, err := cu.customerRepository.GetCustomer(id)
	if err != nil {
		return err
	}

	err = cu.customerRepository.Delete(*customer)
	if err != nil {
		return err
	}

	return nil
}
