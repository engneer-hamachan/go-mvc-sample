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

func (cr *customerPersistance) GetCustomer(id int) model.Customer {

	var customer model.Customer
	cr.Conn.First(&customer, id)

	return customer

}

func (cr *customerPersistance) GetCustomers() []model.Customer {

	var customers []model.Customer
	cr.Conn.Find(&customers)

	return customers
}

func (cr *customerPersistance) Create(c model.Customer) {

	cr.Conn.Create(&c)

}

func (cr *customerPersistance) Update(c model.Customer) {

	cr.Conn.Save(&c)

}

func (cr *customerPersistance) Delete(c model.Customer) {

	cr.Conn.Delete(&c)

}
