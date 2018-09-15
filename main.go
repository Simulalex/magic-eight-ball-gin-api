package main

import (
	"github.com/Simulalex/magic-eight-ball-gin-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	handlers := handlers.Create("db.sql")

	r.GET("/fortunes", handlers.Read)
	r.POST("/fortunes", handlers.Create)
	r.PUT("/fortunes/:id", handlers.Update)
	r.DELETE("/fortunes/:id", handlers.Delete)

	r.Run()
}
