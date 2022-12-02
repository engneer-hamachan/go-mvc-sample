package controller

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"strconv"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func StoreCustomer(c *gin.Context) {
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))

	customer := model.Customer{Name: name, Age: age}
	customer.Create()

	c.Redirect(301, "/")
}
