package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	//加载html文件
	engine.LoadHTMLGlob("./html/*")
	//加载img文件
	engine.Static("/img", "./img")
	//http://localhost:8080/helloHtml
	engine.GET("helloHtml", func(context *gin.Context) {
		fullPath := "请求路径" + context.FullPath()
		fmt.Println(fullPath)
		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullPath": fullPath,
			"title":    "Gin教程",
		})
	})
	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
