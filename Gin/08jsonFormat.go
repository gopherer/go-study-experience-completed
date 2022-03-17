package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Response struct {
	Code    int
	Message string
	Data    string
}

func main() {
	engine := gin.Default()
	//http://localhost:8080/jsonMap
	engine.GET("jsonMap", func(context *gin.Context) {
		fullPath := "请求路径" + context.FullPath()
		fmt.Println(fullPath)
		m := map[string]interface{}{
			"code":    1,
			"message": "ok",
		}
		m["data"] = fullPath
		context.JSON(200, &m)
		context.JSON(200, &map[string]interface{}{
			"code":    1,
			"message": "OK",
			"data":    "hello",
		})
	})
	//http://localhost:8080/jsonStruct
	engine.GET("jsonStruct", func(context *gin.Context) {
		fullPath := "请求路径" + context.FullPath()
		fmt.Println(fullPath)
		context.JSON(200, &struct {
			Code    int
			Message string
			Data    string
		}{1, "OK", fullPath})
		res := Response{
			Code:    1,
			Message: "OK",
			Data:    fullPath,
		}
		context.JSON(200, &res)
	})
	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
