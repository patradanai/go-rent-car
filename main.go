package main

import (
	"car-booking/configs"
	"car-booking/routes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitialRouter(c *gin.Engine) {
	c.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, "Root Path")
		return
	})

	r := c.Group("/api/v1/")

	routes.AuthRouter(r.Group("/auth"))
}

func main() {
	router := gin.Default()
	if _, err := configs.ConnectToDB(); err != nil {
		panic(err)
	}

	InitialRouter(router)

	server := &http.Server{
		Addr:           ":3000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
