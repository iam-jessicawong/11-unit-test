package repository

import (
	"unit-test/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id uint) (*models.Product, error) {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(0)
	}

	product := arguments.Get(0).(models.Product)
	return &product, nil
}

func (repository *ProductRepositoryMock) FindAll(role string, userId uint) ([]*models.Product, error) {
	arguments := repository.Mock.Called(role, userId)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(0)
	}

	products := arguments.Get(0).([]*models.Product)
	return products, nil
}
