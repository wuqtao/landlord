package config

import (
	"github.com/BurntSushi/toml"
	"fmt"
)

//订制配置文件解析载体
type Config struct{
	Database *Database
}

type Database struct {
	Host string
	Port int
	DbName  string
	Username  string
	Password string
}

var Con *Config=new (Config)

func init(){
	//读取配置文件
	_, err := toml.DecodeFile("E:\\workspace\\go\\GOPATH\\src\\chessSever\\config\\config.toml",Con)
	if err!=nil{
		fmt.Println(err)
	}

}
