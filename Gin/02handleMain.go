package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	//http://localhost:8080/hello?name=davie
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)
		name := context.DefaultQuery("name", "hello")
		if _, err := context.Writer.Write([]byte("hello" + name)); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	})
	//http://localhost:8080/login
	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		username := context.PostForm("username")
		password := context.PostForm("password")
		fmt.Println(username, password)
		if _, err := context.Writer.Write([]byte("login," + username)); err != nil {
			log.Fatal(err)
		}
	})
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}
