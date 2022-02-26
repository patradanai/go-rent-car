package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginInfo := loginBody{}
		if err := c.ShouldBindJSON(&loginInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}

	}
}
