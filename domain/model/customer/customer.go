package customer

import (
	"github.com/google/uuid"
	"main/domain/model/vo"
)

type Customer struct {
	customerId vo.UuId
	name       vo.PersonName
	age        vo.Age
}

func New(customerId string, name string, age int) (*Customer, error) {
	createdCustomerId, err := vo.NewUuId(customerId)
	if err != nil {
		return nil, err
	}

	createdName, err := vo.NewName(name)
	if err != nil {
		return nil, err
	}

	createdAge, err := vo.NewAge(age)
	if err != nil {
		return nil, err
	}

	customer := Customer{
		customerId: *createdCustomerId,
		name:       *createdName,
		age:        *createdAge,
	}

	return &customer, nil
}

// Create Constructor
func Create(name string, age int) (*Customer, error) {
	customerId := uuid.New().String()
	customer, err := New(customerId, name, age)

	if err != nil {
		return nil, err
	}

	return customer, err
}

// Getter
func (c Customer) GetCustomerId() string {
	return string(c.customerId)
}

func (c Customer) GetName() string {
	return string(c.name)
}

func (c Customer) GetAge() int {
	return int(c.age)
}
