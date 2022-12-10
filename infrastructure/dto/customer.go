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
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func ConvertCustomer(c *customer.Customer) *Customer {
	return &Customer{
		CustomerId: string(c.GetCustomerId()),
		Name:       string(c.GetName()),
		Age:        int(c.GetAge()),
	}
}

func AdaptCustomer(converted_customer *Customer) (*customer.Customer, error) {
	customer, err := customer.New(
		converted_customer.CustomerId,
		converted_customer.Name,
		converted_customer.Age,
	)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func AdaptCustomers(converted_customers []*Customer) ([]customer.Customer, error) {
	var customers []customer.Customer

	for _, converted_customer := range converted_customers {
		customer, err := customer.New(
			converted_customer.CustomerId,
			converted_customer.Name,
			converted_customer.Age,
		)

		if err != nil {
			return nil, err
		}
		customers = append(customers, *customer)
	}

	return customers, nil
}
