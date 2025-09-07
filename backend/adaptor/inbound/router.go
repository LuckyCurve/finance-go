package inbound

import "github.com/gin-gonic/gin"

func Run() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

	}
	router.Run("localhost:8081")
}
