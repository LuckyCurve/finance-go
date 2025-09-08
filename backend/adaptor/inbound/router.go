package inbound

import (
	"github.com/gin-gonic/gin"
)

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
	{
		api.GET("/asset/list", func(c *gin.Context) {
			res, err := AssetList()
			JSONWithData(c, res, err)
		})
		api.GET("/asset/listWithExchangeRate", func(c *gin.Context) {
			res, err := AssetListWithExchangeRate(c)
			JSONWithData(c, res, err)
		})
		api.POST("/asset/saveOrUpdate", func(c *gin.Context) {
			err := AssetCreateOrUpdate(c)
			JSON(c, err)
		})
	}

	router.Run("localhost:8081")
}

func JSONWithData(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(200, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}

func JSON(c *gin.Context, err error) {
	if err != nil {
		c.JSON(200, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
	})
}
