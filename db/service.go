package db

import "github.com/0xcregis/easynode_monitor/common"

type SqlInterface interface {
	//user
	Query() ([]*common.SysUser, error)
	AddUser(u **common.SysUser) error
	ResetPwd(account string, pwd string) error
}
