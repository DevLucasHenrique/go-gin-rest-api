package main

import (
	"github.com/DevLucasHenrique/go-gin-rest-api/controller"
	"github.com/DevLucasHenrique/go-gin-rest-api/db"
	"github.com/DevLucasHenrique/go-gin-rest-api/repository"
	"github.com/DevLucasHenrique/go-gin-rest-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/pings", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// GET
	server.GET("/products", ProductController.GetProducts)
	server.GET("/products/:productId", ProductController.GetProductById)

	// PUT
	server.PUT("/products/:productId", ProductController.UpdateProduct)

	// POST
	server.POST("/products", ProductController.CreateProduct)

	server.Run(":8000")
}

/* func main() {
	server := gin.Default()
	server.GET("/message", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, products)
	})

	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, products)
	})

	server.POST("/products", func(ctx *gin.Context) {
		var newProduct Product

		if err := ctx.ShouldBindJSON(&newProduct); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		products = append(products, newProduct)

		ctx.JSON(http.StatusCreated, newProduct)
	})

	server.Run(":8000")
} */

// MODEL
/*
type Product struct {
	ID uint `json:id`
	Name string `json:name`
	Price float64 `json:price`
}

// Repository
type ProductRepository struct { // defini a struct do productRepository
	products []Product // falei que os produtos vão ser uma array do model product
}

// r = reciver, aqui estou passando essa função como método do ProductRepository, falando que ela retorna os r.Produtos
func (r *ProductRepository) GetAll() []Product {
	return r.products
}

// Passando outro metodo para o ProductRepository falando que vai receber um product do model Product e vai retornar um product
func (r *ProductRepository) Save(product Product) Product {
	r.products = append(r.products, product) // falando que a array products vai receber o product que foi passado como parametro
	return product
}

// USECASES/SERVICE

type ProductService struct { // defini


	repo *ProductRepository
}

func (s *ProductService) ListProducts() []Product {
	return s.repo.GetAll()
}

func (s *ProductService) CreateProduct(product Product) (Product, error) {
	if product.Price <= 0 {
		return Product{}, errors.New("Product Price cannot be 0 or less than 0")
	}
	if len(product.Name) == 0 {
		return Product{}, errors.New("Name cannot be nil")
	}
	s.repo.Save(product)
	return product, nil
}

// HANDLER/CONTROLLERS
type ProductHandler struct {
	service *ProductService
}

func (service *ProductHandler) GetProducts(ctx *gin.Context) {
	products := service.service.ListProducts()
	ctx.JSON(http.StatusOK, products)
}

func (service *ProductHandler) CreateProduct(ctx *gin.Context) {
	var product Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Dados Invalidos",
		})
		return
	}
	newProduct, err := service.service.CreateProduct(product)
	if err!=nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, newProduct)
}

func main() {
	repo := &ProductRepository{
		products: []Product{{ID: 1, Name: "Teclado Mech", Price: 250.0}},
	}
	service := &ProductService{repo: repo}
	handler := &ProductHandler{service: service}

	server := gin.Default()

	server.GET("/products", handler.GetProducts)
	server.POST("/products", handler.CreateProduct)

	server.Run(":8000")
}

*/
