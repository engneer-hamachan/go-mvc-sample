package persistance

import (
	"gorm.io/gorm"
	"main/domain/model/customer"
	"main/domain/repository"
	"main/infrastructure/dto"
)

type customerPersistance struct {
	Conn *gorm.DB
}

func NewCustomerPersistance(conn *gorm.DB) repository.CustomerRepository {
	return &customerPersistance{Conn: conn}
}

func (cp *customerPersistance) GetCustomerByEmail(email string) (result *customer.Customer, err error) {

	var customer dto.Customer
	if result := cp.Conn.Where("email = ?", email).First(&customer); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_customer, err := dto.AdaptCustomer(&customer)
	if err != nil {
		return nil, err
	}

	return result_customer, nil
}

func (cp *customerPersistance) GetCustomer(id string) (result *customer.Customer, err error) {

	var customer dto.Customer
	if result := cp.Conn.Where("customer_id = ?", id).First(&customer); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_customer, err := dto.AdaptCustomer(&customer)
	if err != nil {
		return nil, err
	}

	return result_customer, nil
}

func (cp *customerPersistance) GetCustomers() (result []*customer.Customer, err error) {

	var customers []*dto.Customer
	if result := cp.Conn.Find(&customers); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_customers, err := dto.AdaptCustomers(customers)
	if err != nil {
		return nil, err
	}

	return result_customers, nil
}

func (cp *customerPersistance) InsertCustomer(c *customer.Customer) error {
	converted_customer := dto.ConvertCustomer(c)

	if result := cp.Conn.Create(converted_customer); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (cp *customerPersistance) UpdateCustomer(c *customer.Customer) error {
	converted_customer := dto.ConvertCustomer(c)

	if result := cp.Conn.Where("customer_id = ?", converted_customer.CustomerId).
		Updates(converted_customer); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (cp *customerPersistance) DeleteCustomer(c *customer.Customer) error {
	converted_customer := dto.ConvertCustomer(c)

	if result := cp.Conn.Where("customer_id = ?", converted_customer.CustomerId).
		Delete(converted_customer); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}
