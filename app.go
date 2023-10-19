package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/0xcregis/easynode_monitor/common/config"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sunjiangjun/xlog"
)

var identityKey = "id"

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

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "easynode",
		Key:         []byte("abcABC1234567890"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims[identityKey].(string)
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			account, err := h.Login(c)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			} else {
				return account, nil
			}
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	g := e.Group(cfg.RootPath)
	g.POST("/login", authMiddleware.LoginHandler)
	u := g.Group("/sys/user")
	u.Use(h.Authenticate)
	u.GET("/list", h.Query)
	u.POST("/add", h.AddUser)
	u.POST("/resetPwd", h.ResetPwd)

	err = e.Run(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		panic(err)
	}
}
