package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
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

	//r.GET("/", func(c *gin.Context) {
	// 传递参数:
	r.GET("/:id", func(c *gin.Context) {
		//c.String(200, "Hello World")
		// 转换为 int
		id, _ := strconv.Atoi(c.Param("id"))
		//id := c.Param("id")

		c.JSON(200, gin.H{
			//"mess": "Hello World",
			"uid":  id,
			"user": user,
		})
	})

	r.Run(":910")
}
