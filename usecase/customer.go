package usecase

import (
	"main/domain/model/customer"
	"main/domain/repository"
)

type CustomerUseCase interface {
	GetCustomer(id string) (result *customer.Customer, err error)
	GetCustomers() (result []customer.Customer, err error)
	CreateCustomer(name string, age int) error
	UpdateCustomer(id string, name string, age int) error
	DeleteCustomer(id string) error
}

type customerUseCase struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerUseCase(cr repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepository: cr,
	}
}

func (cu *customerUseCase) GetCustomer(id string) (result *customer.Customer, err error) {
	customer, err := cu.customerRepository.GetCustomer(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (cu *customerUseCase) GetCustomers() (result []customer.Customer, err error) {
	customers, err := cu.customerRepository.GetCustomers()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (cu *customerUseCase) CreateCustomer(name string, age int) error {
	customer, err := customer.Create(name, age)
	if err != nil {
		return err
	}

	err = cu.customerRepository.Create(customer)
	if err != nil {
		return err
	}

	return nil
}

func (cu *customerUseCase) UpdateCustomer(id string, name string, age int) error {
	current_customer, err := cu.customerRepository.GetCustomer(id)
	if err != nil {
		return err
	}

	customerId := string(current_customer.GetCustomerId())

	update_customer, err := customer.New(customerId, name, age)
	if err != nil {
		return err
	}
	err = cu.customerRepository.Update(update_customer)
	if err != nil {
		return err
	}

	return nil
}

func (cu *customerUseCase) DeleteCustomer(id string) error {
	customer, err := cu.customerRepository.GetCustomer(id)
	if err != nil {
		return err
	}

	err = cu.customerRepository.Delete(customer)
	if err != nil {
		return err
	}

	return nil
}
