package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Register struct {
	Username string `form:"username"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

func main() {
	engine := gin.Default()

	//http://localhost:8080/register
	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var register Register
		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(register.Username)
		fmt.Println(register.Phone)
		if _, err := context.Writer.Write([]byte(register.Username + " register")); err != nil {
			log.Fatal(err.Error())
		}
	})
	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
