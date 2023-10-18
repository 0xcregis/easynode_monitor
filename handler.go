package main

import (
	"github.com/0xcregis/easynode_monitor/common/config"
	"github.com/0xcregis/easynode_monitor/service"
	"github.com/gin-gonic/gin"
	"github.com/sunjiangjun/xlog"
)

type Handler struct {
	userClient service.SysUserInterface
	cfg        *config.Config
	log        *xlog.XLog
}

func NewHandler(cfg *config.Config, x *xlog.XLog) Handler {
	u := service.NewUserSrv(x, cfg)
	return Handler{
		userClient: u,
		cfg:        cfg,
		log:        x,
	}
}

func (h *Handler) Query(ctx *gin.Context) {

}

func (h *Handler) AddUser(ctx *gin.Context) {

}

func (h *Handler) ResetPwd(ctx *gin.Context) {

}
