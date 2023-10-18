package db

import (
	"github.com/0xcregis/easynode_monitor/common"
	"github.com/0xcregis/easynode_monitor/common/config"
	"github.com/0xcregis/easynode_monitor/common/driver"
	"github.com/sunjiangjun/xlog"
	"gorm.io/gorm"
)

type MysqlServer struct {
	db  *gorm.DB
	log *xlog.XLog
}

func NewMysql(x *xlog.XLog, db *config.BaseDb) SqlInterface {
	b, err := driver.Open(db.User, db.Password, db.Addr, db.DbName, db.Port, x)
	if err != nil {
		panic(err)
	}

	return &MysqlServer{
		db:  b,
		log: x,
	}
}

func (m *MysqlServer) Query() ([]*common.SysUser, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MysqlServer) AddUser(u **common.SysUser) error {
	//TODO implement me
	panic("implement me")
}

func (m *MysqlServer) ResetPwd(account string, pwd string) error {
	//TODO implement me
	panic("implement me")
}
