package main

import (
	"main.go/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	gin.SetMode(gin.ReleaseMode)
	// Routes
	r.POST("/payments", controllers.CreatePayment)

	// Start server
	r.Run(":8080")
}
