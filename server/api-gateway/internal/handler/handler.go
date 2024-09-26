package handler

import (
	"gateway/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(config config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.CORS.AllowedOrigins,
		AllowMethods:     config.CORS.AllowedMethods,
		AllowHeaders:     config.CORS.AllowedHeaders,
		ExposeHeaders:    config.CORS.ExposedHeaders,
		AllowCredentials: config.CORS.AllowCredentials,
	}))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	return router
}
