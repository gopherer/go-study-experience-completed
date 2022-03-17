package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	userGroup := engine.Group("/user")

	userGroup.POST("/register", registerHandle)
	userGroup.POST("/login", loginHandle)
	userGroup.DELETE("/:id", deleteHandle)

	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
func registerHandle(context *gin.Context) {
	fullPath := "用户注册" + context.FullPath()
	fmt.Println(fullPath)
	if _, err := context.Writer.WriteString(fullPath); err != nil {
		log.Fatal(err)
	}
}
func loginHandle(context *gin.Context) {
	fullPath := "用户登入" + context.FullPath()
	fmt.Println(fullPath)
	if _, err := context.Writer.WriteString(fullPath); err != nil {
		log.Fatal(err)
	}
}
func deleteHandle(context *gin.Context) {
	fullPath := context.FullPath()
	userId := context.Param("id")
	fmt.Println(fullPath + " " + userId)
	if _, err := context.Writer.WriteString("删除用户" + userId); err != nil {
		log.Fatal(err)
	}
}
