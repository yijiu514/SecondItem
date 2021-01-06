package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"newtest/pkg/models"
	"newtest/pkg/routers"

	"github.com/BurntSushi/toml"

	"github.com/sirupsen/logrus"
)

func TomlRead(path string, v interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("io read wrong %x", err)
	}

	_, err = toml.Decode(string(content), v)
	if err != nil {
		return fmt.Errorf("toml read wrong %x", err)
	}

	return nil
}

var config = flag.String("config", "etc/config.tmol", "config file")

func ItemInit() {

	var db models.TomlConfig
	err := TomlRead(*config, &db)
	if err != nil {
		logrus.WithError(err).Warn("toml read wrong")
	}

	err = models.DBInit(db.DB)
	if err != nil {
		logrus.WithError(err).Warn("DB init wrong")
	}

	err = routers.RouterInit(db.WebPort)
	if err != nil {
		logrus.WithError(err).Warn("Router init wrong")
	}
}
