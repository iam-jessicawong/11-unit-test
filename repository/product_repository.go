package repository

import (
	"unit-test/database"
	"unit-test/models"
)

type ProductRepository interface {
	FindById(id uint) (*models.Product, error)
	FindAll(role string) (*[]models.Product, error)
}

func FindById(id uint) (*models.Product, error) {
	db := database.GetDB()
	product := models.Product{}
	err := db.Debug().Preload("User").First(&product, "id = ?", id).Error

	return &product, err
}
