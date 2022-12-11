package customer

import (
	"fmt"
	"github.com/google/uuid"
)

type Customer struct {
	customerId customerId
	name       name
	age        age
}

type customerId string
type name string
type age int

func New(customerId string, name string, age int) (*Customer, error) {
	createdCustomerId, err := newCustomerId(customerId)
	if err != nil {
		return nil, err
	}

	createdName, err := newName(name)
	if err != nil {
		return nil, err
	}

	createdAge, err := newAge(age)
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

// value constructors
func newCustomerId(value string) (*customerId, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:customerId NewCustomerId()")
		return nil, err
	}

	customerId := customerId(value)

	return &customerId, nil
}

func newName(value string) (*name, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:name newName()")
		return nil, err
	}

	name := name(value)

	return &name, nil
}

func newAge(value int) (*age, error) {
	if value >= 100 {
		err := fmt.Errorf("%s", "invalid age arg:age newAge()")
		return nil, err
	}

	age := age(value)

	return &age, nil
}
