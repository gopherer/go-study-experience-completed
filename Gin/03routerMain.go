package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	//http://localhost:8080/hello?name=davie
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		name := context.Query("name")
		if _, err := context.Writer.Write([]byte("hello," + name)); err != nil {
			log.Fatal(err)
		}
	})
	//http://localhost:8080/login
	engine.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		if username, exist := context.GetPostForm("username"); exist {
			fmt.Println(username)
		}
		if password, exist := context.GetPostForm("password"); exist {
			fmt.Println(password)
		}
		if _, err := context.Writer.Write([]byte("hello")); err != nil {
			log.Fatal(err)
		}
	})
	//http://localhost:8080/user
	engine.DELETE("/user/:id", func(context *gin.Context) {
		userId := context.Param("id")
		fmt.Println(userId)
		if _, err := context.Writer.Write([]byte("delete," + userId)); err != nil {
			log.Fatal(err)
		}
	})
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}
