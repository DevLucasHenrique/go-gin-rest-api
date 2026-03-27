package main

import (
	"go-gin-rest-api/controller"
	"go-gin-rest-api/db"
	"go-gin-rest-api/repository"
	"go-gin-rest-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if(err!=nil) {
		panic(err)
	}

	// Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
