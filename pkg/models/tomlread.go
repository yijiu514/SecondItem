package models

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

// TomlConfig 定义配置文件结构体
type TomlConfig struct {
	WebPort int64
	DB      Database `toml:"database"`
}

// Database 定义数据库信息结构体
type Database struct {
	UserName string
	PassWord string
	IP       string
	Port     string
	DBName   string
}

// TomlRead 读入toml配置文件
func TomlRead() (int64, Database, error) {

	var config TomlConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		logrus.WithError(err).Warn("toml file read failed")
		return 0, Database{}, err
	}

	return config.WebPort, config.DB, nil
}
