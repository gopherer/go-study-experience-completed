package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tools"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"os"
)

func main() {
	cfg, err := tools.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	app := gin.Default()
	if _, err := tools.OrmEngine(cfg); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	registerRouter(app)

	if err := app.Run(cfg.AppHost + ":" + cfg.AppPort); err != nil {
		panic(err.Error())
	}

}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
