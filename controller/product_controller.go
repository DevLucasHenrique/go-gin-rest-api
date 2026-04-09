package controller

import (
	"net/http"
	"strconv"

	"github.com/DevLucasHenrique/go-gin-rest-api/model"
	"github.com/DevLucasHenrique/go-gin-rest-api/usecase"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if product == (model.Product{}) {
		ctx.JSON(http.StatusNoContent, gin.H{
			"ERROR": "No content",
		})
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")

	if id == "" {
		response := model.Response{
			Message: "Id cannot be nothing",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ERROR: Product Id is not a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(uint(productId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if err == nil && product == nil {
		response := model.Response{
			Message: "Product not found in Database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {

	var product model.Product
	product_id := ctx.Param("productId")

	if product_id == "" {
		response := model.Response{
			Message: "Id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(product_id)

	if err != nil {
		response := model.Response{
			Message: "Id need to be a Number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = ctx.BindJSON(&product)

	if (product == model.Product{}) {
		response := model.Response{
			Message: "Product cannot be nil",
		}
		ctx.JSON(http.StatusNoContent, response)
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	product, err = p.productUseCase.UpdateProduct(uint(productId), product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
