package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		//c.String(200, "Hello World")
		c.JSON(200, gin.H{
			"mess": "Hello World",
		})
	})

	r.Run(":910")
}
