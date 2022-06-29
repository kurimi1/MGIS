package svc

import (
	"LGIS/api/internal/config"
	"LGIS/dao"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Dao    *dao.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Db))
	if err != nil {
		logx.Errorf("db error: %v", err)
		os.Exit(1)
	}
	cl := dao.Use(db)
	return &ServiceContext{
		Config: c,
		Dao:    cl,
	}
}
