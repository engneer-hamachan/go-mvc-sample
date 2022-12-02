package main

import (
	"github.com/gin-gonic/gin"
	"main/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*html")

	router.GET("/", controller.Index)
	router.POST("/create", controller.StoreCustomer)

	router.Run(":8080")
}
