package main

import (
	"github.com/gin-gonic/gin"
  "main/controller"
  "main/infrastructure/persistance"
  sqlite "main/config/database"
)

func main() {
  db := sqlite.New()
  connect, _ := db.DB()
  defer connect.Close()

  //DI
  customerRepository := persistance.NewCustomerPersistance(db)
  customerController := controller.NewCustomerController(customerRepository)


	router := gin.Default()
	router.LoadHTMLGlob("view/*html")

	router.GET("/", customerController.Index)
	router.GET("/customer/:id", customerController.DetailCustomer)
	router.POST("/customer/create", customerController.CreateCustomer)
	router.POST("/customer/update", customerController.UpdateCustomer)
	router.POST("/customer/delete", customerController.DeleteCustomer)

	router.Run(":8080")
}
