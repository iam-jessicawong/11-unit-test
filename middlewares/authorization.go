package middlewares

import (
	"net/http"
	"strconv"
	"unit-test/database"
	"unit-test/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid id",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "Product not exist",
			})
			return
		}

		if userData["role"] != "admin" && Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to view or edit this product",
			})
			return
		}

		c.Next()
	}
}

func DeleteAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)

		if userData["role"] != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to delete product",
			})
			return
		}

		c.Next()
	}
}
