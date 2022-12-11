package dto

import (
	"main/domain/model/customer"
	"time"
)

type Customer struct {
	ID         int
	CustomerId string
	Name       string
	Age        int
	Email      string
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func ConvertCustomer(customer *customer.Customer) *Customer {
	return &Customer{
		CustomerId: customer.GetCustomerId(),
		Name:       customer.GetName(),
		Age:        customer.GetAge(),
		Email:      customer.GetEmail(),
		Password:   customer.GetPassword(),
	}
}

func AdaptCustomer(converted_customer *Customer) (*customer.Customer, error) {
	customer, err := customer.New(
		converted_customer.CustomerId,
		converted_customer.Name,
		converted_customer.Age,
		converted_customer.Email,
		converted_customer.Password,
	)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func AdaptCustomers(converted_customers []*Customer) ([]*customer.Customer, error) {
	var customers []*customer.Customer

	for _, converted_customer := range converted_customers {
		customer, err := customer.New(
			converted_customer.CustomerId,
			converted_customer.Name,
			converted_customer.Age,
			converted_customer.Email,
			converted_customer.Password,
		)

		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
