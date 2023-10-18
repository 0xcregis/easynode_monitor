package config

type Config struct {
	RootPath string  `json:"RootPath"`
	Port     int     `json:"Port"`
	LogLevel int     `json:"LogLevel"`
	BaseDb   *BaseDb `json:"BaseDb"`
}

/*
*

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
