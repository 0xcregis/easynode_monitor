package service

import "github.com/0xcregis/easynode_monitor/common"

type SysUserInterface interface {
	Query() ([]*common.SysUser, error)
	GetSysUser(account string) (*common.SysUser, error)
	AddUser(u *common.SysUser) error
	ResetPwd(account string, pwd string) error
}
type ChainNodeInterface interface {
	Query() (*common.ChainNode, error)
	AddFullNode(fullNode *common.FullNode) error
	QueryFullNode(chain string) ([]*common.FullNode, error)
	GetAdvanceFullNode(chain string) (*common.FullNode, error)
	UpdateFullNode(chain string, fullNode *common.FullNode) error
}
