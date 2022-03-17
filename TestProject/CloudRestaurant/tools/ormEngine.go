package tools

import (
	"CloudRestaurant/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/wonderivan/logger"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.CharSet
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	engine.ShowSQL(database.ShowSql)
	if err := engine.Sync2(new(model.SmsCode)); err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm

	return orm, nil
}
