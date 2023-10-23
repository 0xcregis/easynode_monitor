package db

import "github.com/0xcregis/easynode_monitor/common"

type SqlInterface interface {
	//user
	QueryAllSysUser() ([]*common.SysUser, error)
	GetSysUser(account string) (*common.SysUser, error)
	AddSysUser(u *common.SysUser) error
	ResetPwd(account string, pwd string) error

	//fullnode
	AddFullNode(fullNode *common.FullNode) error
	QueryFullNode(chain string) ([]*common.FullNode, error)
	UpdateFullNode(chain string, fullNode *common.FullNode) error
}
