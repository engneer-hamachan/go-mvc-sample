package persistance

import (
	"gorm.io/gorm"
	"main/domain/model"
	"main/domain/repository"
)

type customerPersistance struct {
	Conn *gorm.DB
}

func NewCustomerPersistance(conn *gorm.DB, c repository.CustomerRepository) *customerPersistance {
	return &customerPersistance{Conn: conn}
}

func (cr *customerPersistance) GetCustomer(id int) (result *model.Customer, err error) {

	var customer model.Customer
	if result := cr.Conn.First(&customer, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &customer, nil
}

func (cr *customerPersistance) GetCustomers() (result []model.Customer, err error) {

	var customers []model.Customer
	if result := cr.Conn.Find(&customers); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return customers, nil
}

func (cr *customerPersistance) Create(c model.Customer) error {

	if result := cr.Conn.Create(&c); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (cr *customerPersistance) Update(c model.Customer) error {

	if result := cr.Conn.Save(&c); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (cr *customerPersistance) Delete(c model.Customer) error {

	if result := cr.Conn.Delete(&c); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}
