package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

//`json:"name"`亦可
type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  int    `form:"age"`
}

func main() {
	engine := gin.Default()

	engine.POST("/addStudent", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var person Person
		if err := context.BindJSON(&person); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(person.Name)
		fmt.Println(person.Age)
		if _, err := context.Writer.Write([]byte("添加记录: " + person.Name)); err != nil {
			log.Fatal(err)
		}
	})
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}
