package controller

import (
	"CloudRestaurant/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//主要用于处理用户相关的请求
type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
}

// http://localhost:8080/api/sendcode?phone=123456789
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//发送验证码
	if _, err := context.Writer.WriteString("该功能有待完善"); err != nil {
		log.Fatal(err)
	}
	phone := "333"
	ms := service.MemberService{}
	isSend := ms.SendCode(phone)
	if isSend {
		context.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"msg":  "发送成功",
		})
	} else {
		context.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"msg":  "发送失败",
		})
	}
	/*
		phone, exist := context.GetQuery("phone")
		if !exist {
			context.JSON(http.StatusOK, map[string]interface{}{
				"code": 0,
				"msg":  "参数解析失败",
			})
			return
		}
		ms := service.MemberService{}
		isSend := ms.SendCode(phone)
		if isSend {
			context.JSON(http.StatusOK, map[string]interface{}{
				"code": 1,
				"msg":  "发送成功",
			})
		} else {
			context.JSON(http.StatusOK, map[string]interface{}{
				"code": 0,
				"msg":  "发送失败",
			})
		}
	*/
}
