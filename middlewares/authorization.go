package middlewares

import (
	"car-booking/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type headerAuth struct {
	Token string `header:"Authorization"`
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerAuth := headerAuth{}
		if err := c.ShouldBindHeader(&headerAuth); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "token can not empty"})
			return
		}

		// Split Token

		token := strings.Split(headerAuth.Token, "Bearer ")
		if len(token) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "token can not empty"})
			return
		}

		jwtService := services.NewJWTService()

		tokenCliam, err := jwtService.ValidateToken(token[1])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "something went wrong"})
			return
		}

		cliams := tokenCliam.Claims.(jwt.MapClaims)

		userId := cliams["userId"]

		c.Set("userId", userId)

		c.Next()
	}
}
