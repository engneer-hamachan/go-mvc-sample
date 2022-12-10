package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/usecase"
	"strconv"
)

type customerController struct {
	customerUseCase usecase.CustomerUseCase
}

func NewCustomerController(cu usecase.CustomerUseCase) customerController {
	return customerController{
		customerUseCase: cu,
	}

}

func (cc *customerController) Index(c *gin.Context) {

	customers, err := cc.customerUseCase.GetCustomers()
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	type ResultDataField struct {
		CustomerId string
		Name       string
		Age        int
	}

	var data []ResultDataField
	for _, customer := range customers {
		customerId := string(customer.GetCustomerId())
		name := string(customer.GetName())
		age := int(customer.GetAge())
		data = append(data, ResultDataField{CustomerId: customerId, Name: name, Age: age})
	}

	c.HTML(200, "index.html", gin.H{"customers": data})
}

func (cc *customerController) DetailCustomer(c *gin.Context) {

	id := c.Param("id")
	customer, err := cc.customerUseCase.GetCustomer(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	type ResultDataField struct {
		CustomerId string
		Name       string
		Age        int
	}

	data := ResultDataField{CustomerId: string(customer.GetCustomerId()), Name: string(customer.GetName()), Age: int(customer.GetAge())}

	c.HTML(200, "detail.html", gin.H{"customer": data})
}

func (cc *customerController) CreateCustomer(c *gin.Context) {
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

	err = cc.customerUseCase.CreateCustomer(name, age)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/")
}

func (cc *customerController) UpdateCustomer(c *gin.Context) {

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

	id := form.ID
	name := form.Name

	age, err := strconv.Atoi(form.Age)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	err = cc.customerUseCase.UpdateCustomer(id, name, age)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/")
}

func (cc *customerController) DeleteCustomer(c *gin.Context) {
	type RequestDataField struct {
		ID string `form:"id" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	id := form.ID

	err := cc.customerUseCase.DeleteCustomer(id)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/")
}
