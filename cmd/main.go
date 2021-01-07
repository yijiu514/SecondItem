package main

import (
	"flag"
	"fmt"
	"newtest/pkg/models"
	"newtest/pkg/routers"

	"github.com/BurntSushi/toml"

	"github.com/sirupsen/logrus"
)

var config = flag.String("config", "config.toml", "config file")

// TomlRead 读取toml文件并解析
func TomlRead(path string) (tc models.TomlConfig, err error) {
	//解析
	flag.Parse()
	path = *config

	//读取toml文件
	_, err = toml.DecodeFile(path, &tc)
	if err != nil {
		return models.TomlConfig{}, fmt.Errorf("toml read wrong %m", err)
	}
	return
}

func main() {

	tc, err := TomlRead(*config)
	if err != nil {
		logrus.WithError(err).Warn("toml read wrong")
	}

	err = models.DBInit(tc.DB)
	if err != nil {
		logrus.WithError(err).Warn("DB init wrong")
	}

	err = routers.RouterInit(tc.WebPort)
	if err != nil {
		logrus.WithError(err).Warn("Router init wrong")
	}
}
