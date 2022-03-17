package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	//http://localhost:8090/helloByte
	engine.GET("/helloByte", func(context *gin.Context) {
		fullPath := "请求路径" + context.FullPath()
		fmt.Println(fullPath)
		if _, err := context.Writer.Write([]byte(fullPath)); err != nil {
			log.Fatal(err)
		}
	})
	//http://localhost:8090/helloString
	engine.GET("/helloString", func(context *gin.Context) {
		fullPath := "请求路径" + context.FullPath()
		fmt.Println(fullPath)
		if _, err := context.Writer.WriteString(fullPath); err != nil {
			log.Fatal(err)
		}
	})
	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err)
	}

}
