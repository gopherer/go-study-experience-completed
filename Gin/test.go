package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		context.String(200, "hello")
	})

	if err := engine.Run(); err != nil {
		log.Fatal(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
