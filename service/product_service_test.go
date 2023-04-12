package service

import (
	"errors"
	"testing"
	"unit-test/models"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := models.Product{
		GormModel:   models.GormModel{ID: 1},
		Title:       "Mouse",
		Description: "Black mouse",
		UserID:      1,
	}

	productRepository.Mock.On("FindById", uint(1)).Return(product, nil)

	result, err := productService.GetOneProduct(uint(1))

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, product.GormModel.ID, result.GormModel.ID, "result has to be '1'")
	assert.Equal(t, product.Title, result.Title, "result has to be 'Mouse'")
	assert.Equal(t, product.Description, result.Description, "result has to be 'Black mouse'")
	assert.Equal(t, product.UserID, result.UserID, "result has to be '1'")
}

func TestProductServiceGetOneProduct_NotFound(t *testing.T) {
	productRepository.Mock.On("FindById", uint(5)).Return(nil, errors.New("product not found"))
	result, err := productService.GetOneProduct(uint(5))

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetAllProduct(t *testing.T) {
	products := []*models.Product{
		{
			GormModel:   models.GormModel{ID: 1},
			Title:       "bantal",
			Description: "bantal guling",
			UserID:      2,
		},
		{
			GormModel:   models.GormModel{ID: 2},
			Title:       "keyboard",
			Description: "mechanical keyboar",
			UserID:      2,
		},
	}

	productRepository.Mock.On("FindAll", "user", uint(2)).Return(products, nil)
	result, err := productService.GetAllProducts("user", uint(2))

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, uint(1), result[0].ID, "result has to be '1'")
	assert.Equal(t, "bantal", result[0].Title, "result has to be 'bantal'")
	assert.Equal(t, "bantal guling", result[0].Description, "result has to be 'bantal guling'")
	assert.Equal(t, uint(2), result[0].UserID, "result has to be '2'")
}

func TestProductServiceGetAllProduct_NotFound(t *testing.T) {
	productRepository.Mock.On("FindAll", "user", uint(5)).Return(nil, errors.New("product not found"))
	result, err := productService.GetAllProducts("user", uint(5))

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}
