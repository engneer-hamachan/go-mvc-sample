package controller

import (
	"github.com/gin-gonic/gin"
	"main/domain/model"
	"main/domain/repository"
	"strconv"
)

type customerController struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerController(cr repository.CustomerRepository) customerController {
	return customerController{
		customerRepository: cr,
	}

}

func (cu *customerController) Index(c *gin.Context) {
	customers := cu.customerRepository.GetCustomers()
	c.HTML(200, "index.html", gin.H{"customers": customers})
}

func (cu *customerController) DetailCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer := cu.customerRepository.GetCustomer(id)
	c.HTML(200, "detail.html", gin.H{"customer": customer})
}

func (cu *customerController) CreateCustomer(c *gin.Context) {
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))

	customer := model.Customer{Name: name, Age: age}
	cu.customerRepository.Create(customer)

	c.Redirect(301, "/")
}

func (cu *customerController) UpdateCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))

	customer := cu.customerRepository.GetCustomer(id)

	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))

	customer.Name = name
	customer.Age = age
	cu.customerRepository.Update(customer)

	c.Redirect(301, "/")
}

func (cu *customerController) DeleteCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))

	customer := cu.customerRepository.GetCustomer(id)

	cu.customerRepository.Delete(customer)

	c.Redirect(301, "/")
}
