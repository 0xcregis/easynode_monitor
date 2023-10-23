package db

import (
	"testing"

	"github.com/0xcregis/easynode_monitor/common"
	"github.com/0xcregis/easynode_monitor/common/config"
	"github.com/sunjiangjun/xlog"
)

func Init() SqlInterface {
	c := config.LoadConfig("../config.json")
	return NewMysql(xlog.NewXLogger(), c.BaseDb)
}
func TestMysqlServer_AddFullNode(t *testing.T) {
	s := Init()
	err := s.AddFullNode(&common.FullNode{
		ChainCode:         "200",
		ChainName:         "eth",
		ChainUri:          "https://eth.merkle.io",
		NodeUri:           "https://eth.merkle.io",
		NodeJwt:           "",
		LatestBlockNumber: "1222",
		NetStatus:         0,
		NetDuration:       1000,
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}

func TestMysqlServer_AddSysUser(t *testing.T) {
	s := Init()
	err := s.AddSysUser(&common.SysUser{
		Account:  "123",
		Password: "123",
		NickName: "admin",
		Role:     1,
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(err)
	}
}

func TestMysqlServer_GetSysUser(t *testing.T) {
	s := Init()
	u, err := s.GetSysUser("123")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(u)
	}
}

func TestMysqlServer_QueryAllSysUser(t *testing.T) {
	s := Init()
	list, err := s.QueryAllSysUser()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(list)
	}
}

func TestMysqlServer_QueryFullNode(t *testing.T) {
	s := Init()
	list, err := s.QueryFullNode("200")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(list)
	}
}

func TestMysqlServer_ResetPwd(t *testing.T) {
	s := Init()
	err := s.ResetPwd("123", "888")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(err)
	}
}

func TestMysqlServer_UpdateFullNode(t *testing.T) {
	s := Init()
	err := s.UpdateFullNode("200", &common.FullNode{
		ChainCode:         "200",
		ChainName:         "eth",
		ChainUri:          "https://eth.merkle.io",
		NodeUri:           "https://eth.merkle.io",
		NodeJwt:           "",
		LatestBlockNumber: "1234",
		NetStatus:         0,
		NetDuration:       1001,
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}
