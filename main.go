package main

import (
	"jobhub/controllers"
	"jobhub/middlewares"
	"jobhub/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	userRoute := r.Group("/user")

	userRoute.POST("/register", controllers.Register)
	userRoute.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
