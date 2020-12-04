package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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

	//  带日志启动，回复，
	// engine.Use(Logger(), Recovery())
	//r := gin.Default()
	//

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Logger()) // 启动日志

	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("./gin/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			//param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery()) // 启动恢复，

	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	/* 登录组件 */
	/*authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{

	}*/

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

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	//r.POST("/upload", func(c *gin.Context) {
	r.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "./1.txt")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	r.Run(":910")
}
