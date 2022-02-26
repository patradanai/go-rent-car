package controllers

import (
	"car-booking/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	userModel := models.User{}

	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	fmt.Println(userModel)
}
