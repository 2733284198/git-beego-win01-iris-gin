package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int64
	Name string
	Age  int
}

func init() {

}

func main() {
	user := User{
		Id:   1,
		Name: "manlan1",
		Age:  20,
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		//c.String(200, "Hello World")
		c.JSON(200, gin.H{
			//"mess": "Hello World",
			"user": user,
		})
	})

	r.Run(":910")
}
