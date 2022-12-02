package main

import (
	"github.com/gin-gonic/gin"
	"main/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*html")

	router.GET("/", controller.Index)
	router.GET("/:id", controller.DetailCustomer)
	router.POST("/create", controller.CreateCustomer)
	router.POST("/update", controller.UpdateCustomer)
	router.POST("/delete", controller.DeleteCustomer)

	router.Run(":8080")
}
