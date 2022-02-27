package middlewares

import (
	"car-booking/configs"
	repository "car-booking/repositories"
	service "car-booking/services"
	"car-booking/utils"
	"fmt"
	"net/http"
	"time"

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

		userRepo := repository.NewUserRepository(configs.DB)
		userService := service.NewUserService(userRepo)

		// Find Username
		userInfo, err := userService.FindUser(loginInfo.Username)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "user not exist in system"})
			return
		}

		// Compare
		if !utils.DecryptPassowrd(userInfo.Password, loginInfo.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "unauthorization"})
			return
		}

		// Token Access
		jwtService := service.NewJWTService()
		token, err := jwtService.GenerateToken(userInfo.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "somthing went wrong"})
			return
		}

		// Token Refresh
		refreshRepo := repository.NewRefreshTokenRepository(configs.DB)
		refreshService := service.NewRefreshTokenService(refreshRepo)

		uuid := utils.GenerateUUID()
		expiredAt := time.Now().AddDate(0, 0, 15)

		// model
		if _, err := refreshService.CreateRefreshToken(userInfo.ID, uuid, expiredAt, false); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "something went wrong"})
			return
		}

		// Add Roles
		roles := make([]string, 0)
		for _, role := range userInfo.Roles {
			roles = append(roles, role.Name)
		}

		data := map[string]interface{}{
			"access-token":  token,
			"refresh-token": uuid,
			"roles":         roles,
			"permission":    "",
		}

		c.JSON(http.StatusAccepted, gin.H{"success": true, "data": data})
	}
}
