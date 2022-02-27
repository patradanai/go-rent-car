package middlewares

import (
	"car-booking/configs"
	"car-booking/repositories"
	"car-booking/services"
	"car-booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Permission(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("userId").(uint)

		// Find User by Id
		userRepo := repositories.NewUserRepository(configs.DB)
		userService := services.NewUserService(userRepo)

		result, err := userService.FindUserById(userId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "not found user"})
			return
		}

		roles := make([]string, 0)
		for _, role := range result.Roles {
			roles = append(roles, role.Name)
		}

		if !utils.Contains(roles, role) {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "access deniend"})
			return
		}

		c.Next()
	}
}
