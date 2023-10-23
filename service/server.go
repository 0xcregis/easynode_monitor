package service

import (
	"context"
	"fmt"
	"time"

	"github.com/0xcregis/easynode_monitor/common"
	"github.com/0xcregis/easynode_monitor/common/config"
	"github.com/0xcregis/easynode_monitor/db"
	"github.com/0xcregis/easynode_monitor/service/chain"
	"github.com/sirupsen/logrus"
	"github.com/sunjiangjun/xlog"
)

type UserServer struct {
	baseDb db.SqlInterface
	log    *logrus.Entry
}

func NewUserSrv(x *xlog.XLog, cfg *config.Config) SysUserInterface {
	sql := db.NewMysql(x, cfg.BaseDb)
	l := x.WithField("model", "sysUser")
	return &UserServer{
		baseDb: sql,
		log:    l,
	}
}

func (s *UserServer) GetSysUser(account string) (*common.SysUser, error) {
	return s.baseDb.GetSysUser(account)
}

func (s *UserServer) QueryAllSysUser() ([]*common.SysUser, error) {
	return s.baseDb.QueryAllSysUser()
}

func (s *UserServer) AddSysUser(u *common.SysUser) error {
	return s.baseDb.AddSysUser(u)
}

func (s *UserServer) ResetPwd(account string, pwd string) error {
	return s.baseDb.ResetPwd(account, pwd)
}

type ChainNode struct {
	baseDb    db.SqlInterface
	log       *logrus.Entry
	chainPath string
	chains    []*config.Chain
}

// GetAdvanceFullNode get advance fullNode
func (c *ChainNode) GetAdvanceFullNode(chain string) (*common.FullNode, error) {
	list, err := c.QueryFullNode(chain)
	if err != nil {
		return nil, err
	}
	temp := list[0]
	for _, v := range list {
		/**
		  what is advance fullNode:
		  1. netStatus is able
		  2. BlockNumber is more than other
		*/
		if v.NetStatus == 0 || v.NetStatus == 1 {
			if /*v.NetDuration >= 0 && v.NetDuration <= temp.NetDuration && */ v.LatestBlockNumber >= temp.LatestBlockNumber {
				temp = v
			}
		}
	}

	return temp, nil
}

func (c *ChainNode) AddFullNode(fullNode *common.FullNode) error {
	return c.baseDb.AddFullNode(fullNode)
}

func (c *ChainNode) QueryFullNode(chain string) ([]*common.FullNode, error) {
	return c.baseDb.QueryFullNode(chain)
}

func (c *ChainNode) UpdateFullNode(chainCode string, fullNode *common.FullNode) error {
	defer func() {
		//update db
		err := c.baseDb.UpdateFullNode(chainCode, fullNode)
		c.log.Warnf("UpdateFullNode|err:%v", err.Error())
	}()
	number, duration, err := chain.GetLatestBlock(chainCode, fullNode)
	if err != nil {
		fullNode.NetDuration = -1
		fullNode.NetStatus = 2
		return err
	}
	fullNode.LatestBlockNumber = number
	fullNode.NetDuration = duration
	if duration < 1000 {
		fullNode.NetStatus = 0
	} else {
		fullNode.NetStatus = 1
	}

	return nil
}

func NewChainNodeSrv(x *xlog.XLog, cfg *config.Config, ctx context.Context) ChainNodeInterface {
	sql := db.NewMysql(x, cfg.BaseDb)
	l := x.WithField("model", "chainNode")
	c := &ChainNode{
		baseDb:    sql,
		log:       l,
		chainPath: cfg.ChainPath,
	}
	c.LoadChains(ctx)
	c.RefreshFullNode(ctx)
	return c
}

func (c *ChainNode) LoadChains(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				list, err := config.LoadChains(c.chainPath)
				if err != nil {
					c.log.Warnf("LoadChains|err:%v", err.Error())
				}
				c.chains = list

			}

			<-time.After(30 * time.Second)
		}
	}()
}

func (c *ChainNode) RefreshFullNode(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				for _, n := range c.chains {
					code := fmt.Sprintf("%v", n.ChainCode)
					list, err := c.QueryFullNode(code)
					if err != nil {
						break
					}

					for _, v := range list {
						go func(chain string, fullNode *common.FullNode) {
							_ = c.UpdateFullNode(code, fullNode)
						}(code, v)
					}
				}

			}

			<-time.After(10 * time.Minute)
		}
	}()
}

func (c *ChainNode) Query(chain string) (*common.ChainNode, error) {
	for _, v := range c.chains {
		if fmt.Sprintf("%v", v.ChainCode) == chain {
			r := &common.ChainNode{
				Chain: v,
			}
			f, err := c.GetAdvanceFullNode(chain)
			if err != nil {
				return nil, err
			}
			r.FullNode = []*common.FullNode{f}
			return r, nil
		}
	}
	return nil, fmt.Errorf("no record")
}
