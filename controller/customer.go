package controller

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"strconv"
)

func Index(c *gin.Context) {
	customers := model.GetCustomers()
	c.HTML(200, "index.html", gin.H{"customers": customers})
}

func DetailCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer := model.GetCustomer(id)
	c.HTML(200, "detail.html", gin.H{"customer": customer})
}

func CreateCustomer(c *gin.Context) {
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))

	customer := model.Customer{Name: name, Age: age}
	customer.Create()

	c.Redirect(301, "/")
}

func UpdateCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))

	customer := model.GetCustomer(id)

	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))

	customer.Name = name
	customer.Age = age
	customer.Update()

	c.Redirect(301, "/")
}

func DeleteCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))

	customer := model.GetCustomer(id)

	customer.Delete()

	c.Redirect(301, "/")
}
