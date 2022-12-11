package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	sqlite "main/config/database"
	"main/controller"
	"main/infrastructure/persistance"
	"main/middleware"
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

	store := cookie.NewStore([]byte("kokonisecretkeyiretene"))
	router.Use(sessions.Sessions("session", store))

	router.GET("/", customerController.Index)
	router.POST("/customer/create", customerController.CreateCustomer)
	router.POST("/customer/update", customerController.UpdateCustomer)
	router.POST("/customer/delete", customerController.DeleteCustomer)

	router.GET("/login", customerController.Login)
	router.POST("/login", customerController.AuthLogin)
	router.GET("/logout", customerController.Logout)

	authUserGroup := router.Group("/")
	authUserGroup.Use(middleware.IsLogin())
	{
		authUserGroup.GET("/customer/:id", customerController.DetailCustomer)
	}

	router.Run(":8080")
}
