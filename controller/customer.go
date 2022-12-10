package controller

import (
	"fmt"
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

	customers, err := cu.customerRepository.GetCustomers()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(200, "index.html", gin.H{"customers": customers})
}

func (cu *customerController) DetailCustomer(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	customer, err := cu.customerRepository.GetCustomer(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(200, "detail.html", gin.H{"customer": *customer})
}

func (cu *customerController) CreateCustomer(c *gin.Context) {
	type RequestDataField struct {
		Name string `form:"name" binding:"required"`
		Age  string `form:"age" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	name := form.Name
	age, err := strconv.Atoi(form.Age)
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	customer := model.Customer{Name: name, Age: age}
	err = cu.customerRepository.Create(customer)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/")
}

func (cu *customerController) UpdateCustomer(c *gin.Context) {

	type RequestDataField struct {
		ID   string `form:"id" binding:"required"`
		Name string `form:"name" binding:"required"`
		Age  string `form:"age" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(form.ID)
	age, err := strconv.Atoi(form.Age)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	name := form.Name

	customer, err := cu.customerRepository.GetCustomer(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	customer.Name = name
	customer.Age = age
	err = cu.customerRepository.Update(*customer)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/")
}

func (cu *customerController) DeleteCustomer(c *gin.Context) {
	type RequestDataField struct {
		ID string `form:"id" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(form.ID)
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	customer, err := cu.customerRepository.GetCustomer(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	err = cu.customerRepository.Delete(*customer)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/")
}
