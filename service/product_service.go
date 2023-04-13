package service

import (
	"errors"
	"unit-test/models"
	"unit-test/repository"
)

type Services interface {
	GetOneProduct(id uint) (*models.Product, error)
	GetAllProducts(role string, userId uint) ([]models.Product, error)
}

type ProductService struct {
	Repository repository.ProductRepository
	PRStruct   repository.ProductRepositoryStruct
}

func (service *ProductService) GetOneProduct(id uint) (*models.Product, error) {
	product, err := service.PRStruct.FindById(id)

	if err != nil {
		return nil, err
	}

	if product == nil {
		return product, errors.New("product not found")
	}

	return product, err
}

func (service *ProductService) GetAllProducts(role string, userId uint) ([]models.Product, error) {
	products, err := service.PRStruct.FindAll(role, userId)

	if err != nil {
		return nil, err
	}

	if products == nil {
		return products, errors.New("product not found")
	}

	return products, err
}
