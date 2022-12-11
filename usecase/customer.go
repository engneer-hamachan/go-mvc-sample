package usecase

import (
	"main/domain/model/customer"
	"main/domain/repository"
)

type CustomerUseCase interface {
	GetCustomerForAuth(email string) (result *customer.Customer, err error)
	GetCustomer(id string) (result *customer.Customer, err error)
	GetCustomers() (result []*customer.Customer, err error)
	CreateCustomer(name string, age int, email string, password string) error
	UpdateCustomer(id string, name string, age int, email string, password string) error
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

func (cu *customerUseCase) GetCustomerForAuth(email string) (result *customer.Customer, err error) {
	current_customer, err := cu.customerRepository.GetCustomerByEmail(email)
	if err != nil {
		return nil, err
	}

	return current_customer, nil
}

func (cu *customerUseCase) GetCustomers() (result []*customer.Customer, err error) {
	customers, err := cu.customerRepository.GetCustomers()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (cu *customerUseCase) CreateCustomer(name string, age int, email string, password string) error {
	customer, err := customer.Create(name, age, email, password)
	if err != nil {
		return err
	}

	err = cu.customerRepository.InsertCustomer(customer)
	if err != nil {
		return err
	}

	return nil
}

func (cu *customerUseCase) UpdateCustomer(id string, name string, age int, email string, password string) error {
	current_customer, err := cu.customerRepository.GetCustomer(id)
	if err != nil {
		return err
	}

	customerId := current_customer.GetCustomerId()

	update_customer, err := customer.New(customerId, name, age, email, password)
	if err != nil {
		return err
	}
	err = cu.customerRepository.UpdateCustomer(update_customer)
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

	err = cu.customerRepository.DeleteCustomer(customer)
	if err != nil {
		return err
	}

	return nil
}
