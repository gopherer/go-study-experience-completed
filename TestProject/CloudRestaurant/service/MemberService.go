package service

import (
	"CloudRestaurant/dataAccess"
	"CloudRestaurant/model"
	"CloudRestaurant/tools"
	"fmt"
	"math/rand"
	"time"
)

type MemberService struct {
}

//go get github.con/aliyun/alibaba-cloud-sdk-go  该命令尚未执行 第三方库目前尚未安装
func (ms *MemberService) SendCode(phone string) bool {
	//1.产生验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	smsCode := model.SmsCode{
		Phone:      phone,
		BizId:      "",
		Code:       code,
		CreateTime: time.Now().Unix(),
	}
	memberDao := dataAccess.MemberDao{tools.DbEngine}
	result := memberDao.InsertCode(smsCode)
	fmt.Println(result, code)
	return result > 0
	/*
		//2.调用阿里云sdk 完成发送 以阿里云为例
		config := tools.GetConfig().Sms
		//client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "<accessKey", "<accessSecret")
		client, err := dysmsapi.NewClientWithAccessKey(config.RegionID, config.AppKey, config.AppSecret)
		if err != nil {
			logger.Error(err.Error())
			return false
		}

		request := dysmsapi.CreatSendSmsRequest()
		request.Scheme = "https"
		request.SignName = config.SignName
		request.TemplateCode = config.TemplateCode
		request.PhoneNumber = phone
		par, err := json.Marshal(map[string]interface{}{
			"code": code,
		})
		response, err := client.SendSms(request)
		request.TemplateParam = string(par)
		fmt.Println(response)
		if err != nil {
			logger.Error(err.Error())
			return false
		}
		//3.接受返回结果，判断发送状态
		//短信验证码发送成功
		if response.Code == "OK" {
			//将验证码保存到数据库
			smsCode := model.SmsCode{
				Phone:      phone,
				BizId:      "",
				Code:       code,
				CreateTime: time.Now().Unix(),
			}
			memberDao := dataAccess.MemberDao{tools.DbEngine}
			result := memberDao.InsertCode(smsCode)
			return result > 0
		}

		return false
	*/
}
