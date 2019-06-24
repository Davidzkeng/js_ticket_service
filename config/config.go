package config

import (
	"github.com/spf13/viper"
)

var Cfg *viper.Viper

func LoadConfigFromToml() error {
	Cfg = viper.New()
	//设置配置文件的名字
	Cfg.SetConfigName("config")

	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	Cfg.AddConfigPath("%GOPATH/src/")
	Cfg.AddConfigPath("./config")

	//设置配置文件类型
	Cfg.SetConfigType("toml")

	if err := Cfg.ReadInConfig(); err != nil{
		return  err
	}

	return nil
}