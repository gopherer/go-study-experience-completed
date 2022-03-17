package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Use(RequestInfos()) //全局使用中间件

	engine.GET("/query", func(context *gin.Context) {
		fmt.Println("中间件的使用方法")
		context.JSON(http.StatusNotFound, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})

	//单独使用中间件
	//engine.GET("/query", RequestInfos(), func(context *gin.Context) {
	//	fmt.Println("中间件的使用方法")
	//	context.JSON(http.StatusNotFound, map[string]interface{}{
	//		"code": 1,
	//		"msg":  context.FullPath(),
	//	})
	//})

	engine.GET("/hello", func(context *gin.Context) {
		if _, err := context.Writer.WriteString("hello"); err != nil {
			log.Fatal(err.Error())
		}
	})
	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}

}

//打印请求信息的中间件
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求路径:", path)
		fmt.Println("请求方法:", method)
		fmt.Println("状态码:", context.Writer.Status())

		context.Next() //中断执行 等待路由执行

		fmt.Println("状态码:", context.Writer.Status())
	}
}
