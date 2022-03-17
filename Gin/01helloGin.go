package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径:", context.FullPath())
		if _, err := context.Writer.Write([]byte("hello gin")); err != nil {
			log.Fatal(err)
		}
	})
	if err := r.Run(); err != nil {
		log.Fatal(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
