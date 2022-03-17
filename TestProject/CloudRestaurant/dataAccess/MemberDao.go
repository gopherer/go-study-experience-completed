package dataAccess

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tools"
	"github.com/wonderivan/logger"
)

type MemberDao struct {
	*tools.Orm
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		logger.Error(err.Error())
	}
	return result
}
