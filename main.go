package main

import (
	"example.com/rest-api/db"
	_ "example.com/rest-api/docs" // Ensure you import the docs package
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db.InitDB()
	defer db.DB.Close()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8080")
}
