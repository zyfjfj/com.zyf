package main

import (
	"fmt"
	"log"
	"time"

	"com.zyf/learn/model"

	"github.com/gin-gonic/gin"
)

var _ model.User

func myMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("I am the middleware")
		//c.Next()
		fmt.Println("I am the after middleware")
	}
}

func main() {
	r := gin.Default()
	user := r.Group("/user", myMiddleware())
	user.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"user": "这是user连接"})
	})
	user.GET("/:id", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)

		}()
		c.String(200, "/user/id")

	})
	r.Run(":8000")
}
