package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/0xcregis/easynode_monitor/common/config"
	"github.com/gin-gonic/gin"
	"github.com/sunjiangjun/xlog"
)

func main() {

	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "The system file of config")
	flag.Parse()
	if len(configPath) < 1 {
		panic("can not find config file")
	}
	cfg := config.LoadConfig(configPath)
	if cfg.LogLevel == 0 {
		cfg.LogLevel = 4
	}
	log.Printf("%+v\n", cfg)

	x := xlog.NewXLogger().BuildOutType(xlog.FILE).BuildLevel(xlog.Level(cfg.LogLevel)).BuildFormatter(xlog.FORMAT_JSON).BuildFile("./log/app", 24*time.Hour)

	e := gin.Default()
	h := NewHandler(&cfg, x)
	g := e.Group(cfg.RootPath)
	u := g.Group("/sys/user")
	u.GET("/list", h.Query)
	u.POST("/add", h.AddUser)
	u.POST("/resetPwd", h.ResetPwd)

	err := e.Run(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		panic(err)
	}
}
