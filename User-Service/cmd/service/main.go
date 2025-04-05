package main

import (
	"github.com/Sapiyulla/Orders-Manager/User-Service/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableQuote:     true,
		DisableTimestamp: true,
	})
}

func main() {
	server := gin.Default()

	server.POST("/register", handlers.Register)
	server.POST("/login", handlers.Login)

	server.GET("/users/:id", handlers.GetUserData)

	server.Run(":8000")
}
