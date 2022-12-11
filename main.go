package main

import (
	"github.com/gin-gonic/gin"
	sqlite "main/config/database"
	"main/controller"
	"main/infrastructure/persistance"
	"main/usecase"
)

func main() {
	db := sqlite.New()
	connect, _ := db.DB()
	defer connect.Close()

	//DI
	customerRepository := persistance.NewCustomerPersistance(db)
	customerUseCase := usecase.NewCustomerUseCase(customerRepository)
	customerController := controller.NewCustomerController(customerUseCase)

	router := gin.Default()
	router.LoadHTMLGlob("view/*html")

	router.GET("/", customerController.Index)
	router.GET("/customer/:id", customerController.DetailCustomer)
	router.POST("/customer/create", customerController.CreateCustomer)
	router.POST("/customer/update", customerController.UpdateCustomer)
	router.POST("/customer/delete", customerController.DeleteCustomer)

	router.Run(":8080")
}
