package db

import (
	"fmt"

	"github.com/0xcregis/easynode_monitor/common"
	"github.com/0xcregis/easynode_monitor/common/config"
	"github.com/0xcregis/easynode_monitor/common/driver"
	"github.com/sunjiangjun/xlog"
	"gorm.io/gorm"
)

const (
	chainNodeTable = "chain_node"
	sysUserTable   = "sys_user"
)

type MysqlServer struct {
	db  *gorm.DB
	log *xlog.XLog
}

func (m *MysqlServer) AddFullNode(fullNode *common.FullNode) error {
	return m.db.Table(chainNodeTable).Omit("id,create_time,update_time").Create(fullNode).Error
}

func (m *MysqlServer) QueryFullNode(chain string) ([]*common.FullNode, error) {
	//net status=0 or 1
	var list []*common.FullNode
	err := m.db.Table(chainNodeTable).Where("net_status !=2 and chain_code=?", chain).Scan(&list).Error
	if err != nil {
		return nil, err
	}

	if len(list) < 1 {
		return nil, fmt.Errorf("no record")
	}
	return list, nil
}

func (m *MysqlServer) UpdateFullNode(chain string, fullNode *common.FullNode) error {
	err := m.db.Table(chainNodeTable).Select("latest_blocknumber,net_status,net_duration").Omit("id,create_time,update_time").Where("chain_code=?", chain).UpdateColumns(fullNode).Error
	if err != nil {
		return err
	}
	return nil
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

func (m *MysqlServer) GetSysUser(account string) (*common.SysUser, error) {
	r := common.SysUser{}
	err := m.db.Table(sysUserTable).Where("account=?", account).First(&r).Error
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (m *MysqlServer) QueryAllSysUser() ([]*common.SysUser, error) {
	var list []*common.SysUser
	err := m.db.Table(sysUserTable).Scan(&list).Error
	if err != nil {
		return nil, err
	}

	if len(list) < 1 {
		return nil, fmt.Errorf("no record")
	}
	return list, nil
}

func (m *MysqlServer) AddSysUser(u *common.SysUser) error {
	return m.db.Table(sysUserTable).Omit("id,create_time,update_time").Create(u).Error
}

func (m *MysqlServer) ResetPwd(account string, pwd string) error {
	return m.db.Table(sysUserTable).Omit("id,create_time,update_time").Where("account=?", account).Update("password", pwd).Error
}
