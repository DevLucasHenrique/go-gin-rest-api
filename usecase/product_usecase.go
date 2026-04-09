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

func (pu *ProductUseCase) GetProductById(product_id uint) (*model.Product, error) {
	product, err := pu.repository.GetProductById(product_id)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) UpdateProduct(product_id uint, product model.Product) (model.Product, error) {
	resultProduct, err := pu.repository.UpdateProduct(product_id, product)

	if err != nil {
		return model.Product{}, err
	}

	return resultProduct, nil
}
