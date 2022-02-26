package routes

import (
	"car-booking/controllers"
	"car-booking/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/signup", controllers.CreateUser)
	r.POST("/signin", middlewares.Authentication())
	r.POST("/refresh")
}
