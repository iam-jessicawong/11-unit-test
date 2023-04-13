package repository

import (
	"unit-test/database"
	"unit-test/models"
)

type ProductRepository interface {
	FindById(id uint) (*models.Product, error)
	FindAll(role string, userId uint) ([]models.Product, error)
}

type ProductRepositoryStruct struct{}

func (*ProductRepositoryStruct) FindById(id uint) (*models.Product, error) {
	db := database.GetDB()
	product := models.Product{}
	err := db.Debug().Preload("User").First(&product, "id = ?", id).Error

	return &product, err
}

func (*ProductRepositoryStruct) FindAll(role string, userId uint) ([]models.Product, error) {
	db := database.GetDB()
	products := []models.Product{}
	var err error

	if role == "admin" {
		err = db.Debug().Preload("User").Find(&products).Error
	} else {
		err = db.Debug().Find(&products, "user_id = ?", userId).Error
	}

	return products, err
}
