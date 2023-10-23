package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/0xcregis/easynode_monitor/common"
	"github.com/0xcregis/easynode_monitor/common/config"
	"github.com/0xcregis/easynode_monitor/service"
	"github.com/gin-gonic/gin"
	"github.com/sunjiangjun/xlog"
	"github.com/tidwall/gjson"
)

type Handler struct {
	userClient  service.SysUserInterface
	chainClient service.ChainNodeInterface
	cfg         *config.Config
	log         *xlog.XLog
}

func NewHandler(cfg *config.Config, x *xlog.XLog, ctx context.Context) Handler {
	u := service.NewUserSrv(x, cfg)
	c := service.NewChainNodeSrv(x, cfg, ctx)
	return Handler{
		userClient:  u,
		cfg:         cfg,
		log:         x,
		chainClient: c,
	}
}

func (h *Handler) Authenticate(ctx *gin.Context) bool {
	return true
}
func (h *Handler) Login(ctx *gin.Context) (string, error) {

	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return "", fmt.Errorf("params is error")
	}
	root := gjson.ParseBytes(b)
	account := root.Get("account").String()
	pwd := root.Get("password").String()
	s1 := md5.Sum([]byte(pwd))
	ss := fmt.Sprintf("%x", s1)
	ss = strings.ToLower(ss)

	u, err := h.userClient.GetSysUser(account)
	if err != nil {
		return "", fmt.Errorf("params is error")
	}

	if u.Password != pwd {
		return "", fmt.Errorf("password:%v is wrong", pwd)
	}
	return account, nil
}

func (h *Handler) QuerySysUser(ctx *gin.Context) {
	list, err := h.userClient.QueryAllSysUser()
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, "", list, ctx.Request.RequestURI)

}

func (h *Handler) AddSysUser(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	var u common.SysUser
	err = json.Unmarshal([]byte(b), &u)
	if err != nil {
		h.Error(ctx, string(b), ctx.Request.RequestURI, err.Error())
		return
	}

	err = h.userClient.AddSysUser(&u)
	if err != nil {
		h.Error(ctx, string(b), ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, string(b), nil, ctx.Request.RequestURI)
}

func (h *Handler) ResetPwd(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	root := gjson.ParseBytes(b)

	account := root.Get("account").String()
	pwd := root.Get("password").String()
	newPwd := root.Get("newPassword").String()

	u, err := h.userClient.GetSysUser(account)
	if err != nil {
		h.Error(ctx, string(b), ctx.Request.RequestURI, err.Error())
		return
	}

	if u.Password != pwd {
		h.Error(ctx, string(b), ctx.Request.RequestURI, "the original password is wrong")
		return
	}

	err = h.userClient.ResetPwd(account, newPwd)
	if err != nil {
		h.Error(ctx, string(b), ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, string(b), nil, ctx.Request.RequestURI)
}

const (
	SUCCESS = 0
	FAIL    = 1
)

func (h *Handler) Success(c *gin.Context, req string, resp interface{}, path string) {
	req = strings.Replace(req, "\t", "", -1)
	req = strings.Replace(req, "\n", "", -1)
	h.log.Printf("path=%v,req=%v,resp=%v\n", path, req, resp)
	mp := make(map[string]interface{})
	mp["code"] = SUCCESS
	mp["data"] = resp
	c.JSON(200, mp)
}

func (h *Handler) Error(c *gin.Context, req string, path string, err string) {
	req = strings.Replace(req, "\t", "", -1)
	req = strings.Replace(req, "\n", "", -1)
	h.log.Errorf("path=%v,req=%v,err=%v\n", path, req, err)
	mp := make(map[string]interface{})
	mp["code"] = FAIL
	mp["data"] = err
	c.JSON(200, mp)
}
