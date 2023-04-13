package controllers

import (
	"log"
	"net/http"
	"strconv"
	"unit-test/database"
	"unit-test/helpers"
	"unit-test/models"
	"unit-test/repository"
	"unit-test/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var productRepository = &repository.ProductRepositoryStruct{}
var productService = service.ProductService{Repository: productRepository}

func GetProducts(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	role := string(userData["role"].(string))
	log.Println(userId, role)
	product, err := productService.GetAllProducts(role, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error getting products data",
			"err":     err.Error(),
		})
		return
	}

	if role == "admin" {
		for _, product := range product {
			product.User.Password = ""
		}

	}

	c.JSON(http.StatusOK, product)
}

func GetProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("id"))
	product, err := productService.GetOneProduct(uint(productId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error getting products data",
			"err":     err.Error(),
		})
		return
	}
	product.User.Password = ""
	c.JSON(http.StatusOK, &product)
}

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Preload("User").Create(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("id"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	productId, _ := strconv.Atoi(c.Param("id"))
	Product := models.Product{}

	err := db.Debug().Where("id = ?", productId).Delete(&Product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Can't delete product",
		})
		return
	}
	c.JSON(http.StatusOK, "Product successfully deleted")
}
