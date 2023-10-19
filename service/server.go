package service

import (
	"github.com/0xcregis/easynode_monitor/common"
	"github.com/0xcregis/easynode_monitor/common/config"
	"github.com/0xcregis/easynode_monitor/db"
	"github.com/sunjiangjun/xlog"
)

type UserServer struct {
	baseDb db.SqlInterface
	x      *xlog.XLog
}

func (s *UserServer) GetSysUser(account string) (*common.SysUser, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserSrv(x *xlog.XLog, cfg *config.Config) SysUserInterface {
	sql := db.NewMysql(x, cfg.BaseDb)
	return &UserServer{
		baseDb: sql,
		x:      x,
	}
}

func (s *UserServer) Query() ([]*common.SysUser, error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserServer) AddUser(u *common.SysUser) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserServer) ResetPwd(account string, pwd string) error {
	//TODO implement me
	panic("implement me")
}
