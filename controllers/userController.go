package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"unit-test/database"
	"unit-test/helpers"
	"unit-test/models"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func Register(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	// trim space to avoid user input white space
	User.Email = strings.TrimSpace(User.Email)
	User.Password = strings.TrimSpace(User.Password)
	User.FullName = strings.TrimSpace(User.FullName)

	err := db.Debug().Create(&User).Error
	if err != nil {
		message := err.Error()
		if err.Error() == "duplicated key not allowed" {
			message = "Your email is already registered, please go to login page"
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"email":     User.Email,
		"full_name": User.FullName,
		"message":   "Thank you for joining us ^^",
	})
}

func Login(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": fmt.Sprintf("user %s is not exist", User.Email),
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "wrong password, please recheck your password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.Role)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
