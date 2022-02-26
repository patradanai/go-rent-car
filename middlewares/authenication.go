package middlewares

import (
	"car-booking/configs"
	repository "car-booking/repositories"
	service "car-booking/services"
	"car-booking/utils"
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

		// Extract Body
		if err := c.ShouldBindJSON(&loginInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}

		userRepo := repository.UserRepository(configs.DB)
		userService := service.NewUserService(userRepo)

		userInfo, err := userService.FindUser(loginInfo.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "user not exist in system"})
			return
		}

		// Compare
		if !utils.DecryptPassowrd(userInfo.Password, loginInfo.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "unauthorization"})
			return
		}

		// Token
		jwtService := service.NewJWTService()
		token, err := jwtService.GenerateToken(userInfo.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "somthing went wrong"})
			return
		}

		data := map[string]string{
			"access-token": token,
		}

		c.JSON(http.StatusAccepted, gin.H{"success": true, "data": data})
	}
}
