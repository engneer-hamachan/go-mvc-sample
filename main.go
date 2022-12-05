package main

import (
	"github.com/gin-gonic/gin"
	"main/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*html")

	router.GET("/", controller.Index)
	router.GET("/customer/:id", controller.DetailCustomer)
	router.POST("/customer/create", controller.CreateCustomer)
	router.POST("/customer/update", controller.UpdateCustomer)
	router.POST("/customer/delete", controller.DeleteCustomer)

	router.Run(":8080")
}
