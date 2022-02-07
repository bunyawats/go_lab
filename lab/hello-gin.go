package lab

import "github.com/gin-gonic/gin"

func HelloGin() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(
			200,
			gin.H{
				"message": "hello world",
			},
		)
	})
	router.Run()
}
