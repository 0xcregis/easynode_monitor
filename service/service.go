package service

import "github.com/0xcregis/easynode_monitor/common"

type SysUserInterface interface {
	QueryAllSysUser() ([]*common.SysUser, error)
	GetSysUser(account string) (*common.SysUser, error)
	AddSysUser(u *common.SysUser) error
	ResetPwd(account string, pwd string) error
}
type ChainNodeInterface interface {
	Query(chain string) (*common.ChainNode, error)
	AddFullNode(fullNode *common.FullNode) error
	QueryFullNode(chain string) ([]*common.FullNode, error)
	GetAdvanceFullNode(chain string) (*common.FullNode, error)
	UpdateFullNode(chain string, fullNode *common.FullNode) error
}
