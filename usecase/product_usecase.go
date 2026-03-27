package usecase

import (
	"github.com/DevLucasHenrique/go-gin-rest-api/model"
	"github.com/DevLucasHenrique/go-gin-rest-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
