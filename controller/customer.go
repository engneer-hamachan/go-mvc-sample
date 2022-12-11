package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func (cc *customerController) Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func (cc *customerController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(200, "logout.html", gin.H{})
}

func (cc *customerController) AuthLogin(c *gin.Context) {
	type RequestDataField struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	email := form.Email
	password := form.Password

	customer, err := cc.customerUseCase.GetCustomerForAuth(email)
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.GetPassword()), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		c.HTML(401, "login.html", gin.H{})
		return
	}

	session := sessions.Default(c)
	session.Set("CustomerId", customer.GetCustomerId())
	session.Save()
	c.Redirect(301, "/")
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
		customerId := customer.GetCustomerId()
		name := customer.GetName()
		age := customer.GetAge()
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
		Email      string
	}

	data := ResultDataField{
		CustomerId: customer.GetCustomerId(),
		Name:       customer.GetName(),
		Age:        customer.GetAge(),
		Email:      customer.GetEmail(),
	}

	c.HTML(200, "detail.html", gin.H{"customer": data})
}

func (cc *customerController) CreateCustomer(c *gin.Context) {
	type RequestDataField struct {
		Name     string `form:"name" binding:"required"`
		Age      string `form:"age" binding:"required"`
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 12)

	name := form.Name
	email := form.Email
	password := string(hash)
	age, err := strconv.Atoi(form.Age)
	if err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	err = cc.customerUseCase.CreateCustomer(name, age, email, password)
	if err != nil {
		fmt.Println(err)
		c.HTML(500, "500.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/")
}

func (cc *customerController) UpdateCustomer(c *gin.Context) {

	type RequestDataField struct {
		ID       string `form:"id" binding:"required"`
		Name     string `form:"name" binding:"required"`
		Age      string `form:"age" binding:"required"`
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var form RequestDataField

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 12)

	id := form.ID
	name := form.Name
	email := form.Email
	password := string(hash)

	age, err := strconv.Atoi(form.Age)
	if err != nil {
		c.HTML(400, "400.html", gin.H{"error": err.Error()})
		return
	}

	err = cc.customerUseCase.UpdateCustomer(id, name, age, email, password)
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
