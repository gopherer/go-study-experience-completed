package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Student struct {
	Name  string `form:"name"`
	Class string `form:"class"`
}

func main() {
	engine := gin.Default()

	//http://localhost:8080/hello?name=davie&class=软件工程
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var student Student
		if err := context.ShouldBindQuery(&student); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(student.Name)
		fmt.Println(student.Class)
		if _, err := context.Writer.Write([]byte("hello," + student.Name)); err != nil {
			log.Fatal(err.Error())
		}
	})
	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
