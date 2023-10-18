package db

import (
	"github.com/0xcregis/easynode_monitor/common/config"
	"github.com/0xcregis/easynode_monitor/common/driver"
	"github.com/sunjiangjun/xlog"
	"gorm.io/gorm"
)

type ChServer struct {
	db  *gorm.DB
	log *xlog.XLog
}

func NewChServer(x *xlog.XLog, db config.BaseDb) SqlInterface {
	b, err := driver.OpenCK(db.User, db.Password, db.Addr, db.DbName, db.Port, x)
	if err != nil {
		panic(err)
	}

	return &MysqlServer{
		db:  b,
		log: x,
	}
}
