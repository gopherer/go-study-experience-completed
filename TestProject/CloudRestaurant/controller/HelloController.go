package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct {
}

func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.Hello)
}

func (hello *HelloController) Hello(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello cloudRestaurant",
	})
}
