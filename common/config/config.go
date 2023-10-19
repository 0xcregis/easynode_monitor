package config

type Config struct {
	RootPath  string  `json:"RootPath"`
	Port      int     `json:"Port"`
	LogLevel  int     `json:"LogLevel"`
	BaseDb    *BaseDb `json:"BaseDb"`
	ChainPath string  `json:"ChainPath"`
}

/*
	{
	    "Addr": "192.168.2.9",
	    "Port": 9000,
	    "User": "test",
	    "Password": "test",
	    "DbName": "base"
	  }
*/
type BaseDb struct {
	User     string `json:"User" gorm:"column:User"`
	Port     int    `json:"Port" gorm:"column:Port"`
	DbName   string `json:"DbName" gorm:"column:DbName"`
	Addr     string `json:"Addr" gorm:"column:Addr"`
	Password string `json:"Password" gorm:"column:Password"`
}

/**
  {
    "chainCode": 200,
    "chainName": "eth",
    "chainIcon": "",
    "httpUri": "",
    "httpPort": "8090",
    "wsUri": "",
    "wsPort": "8091"
  }
*/

type Chain struct {
	ChainName string `json:"chainName" gorm:"column:chainName"`
	ChainIcon string `json:"chainIcon" gorm:"column:chainIcon"`
	WsUri     string `json:"wsUri" gorm:"column:wsUri"`
	HttpUri   string `json:"httpUri" gorm:"column:httpUri"`
	ChainCode int64  `json:"chainCode" gorm:"column:chainCode"`
	HttpPort  string `json:"httpPort" gorm:"column:httpPort"`
	WsPort    string `json:"wsPort" gorm:"column:wsPort"`
}
