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
	//TODO implement me
	panic("implement me")
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

type ChainNode struct {
	baseDb    db.SqlInterface
	log       *logrus.Entry
	chainPath string
	chains    []*config.Chain
}

func (c *ChainNode) GetAdvanceFullNode(chain string) (*common.FullNode, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainNode) AddFullNode(fullNode *common.FullNode) error {
	//TODO implement me
	panic("implement me")
}

func (c *ChainNode) QueryFullNode(chain string) ([]*common.FullNode, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainNode) UpdateFullNode(chainCode string, fullNode *common.FullNode) error {
	defer func() {
		//update db
		err := c.baseDb.UpdateFullNode(chainCode, fullNode)
		c.log.Warnf("UpdateFullNode|err:%v", err.Error())
	}()
	number, duration, err := chain.GetLatestBlock(chainCode)
	if err != nil {
		fullNode.NetDuration = -1
		fullNode.NetStatus = 2
		return err
	}
	fullNode.LatestBlocknumber = number
	fullNode.NetDuration = duration
	if duration < 1000 {
		fullNode.NetStatus = 0
	} else {
		fullNode.NetStatus = 1
	}

	return nil
}

func NewChainNodeSrv(x *xlog.XLog, cfg *config.Config) ChainNodeInterface {
	sql := db.NewMysql(x, cfg.BaseDb)
	l := x.WithField("model", "chainNode")
	return &ChainNode{
		baseDb:    sql,
		log:       l,
		chainPath: cfg.ChainPath,
	}
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

func (c *ChainNode) Query() (*common.ChainNode, error) {
	//TODO implement me
	panic("implement me")
}
