package controllers

import (
	"car-booking/configs"
	"car-booking/models"
	"car-booking/repositories"
	"car-booking/services"
	"car-booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	userModel := models.User{}

	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	userRepo := repositories.NewUserRepository(configs.DB)
	userService := services.NewUserService(userRepo)

	// Encryption Password
	hashedPwd, err := utils.EncryptPassword(userModel.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	userModel.Password = hashedPwd

	if _, err := userService.CreateUser(&userModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "created user success"})
}
