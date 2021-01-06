package controller

import (
	"newtest/pkg/models"
	"newtest/pkg/routers"

	"github.com/sirupsen/logrus"
)

func init() {

	webport, db, err := models.TomlRead()
	if err != nil {
		logrus.WithError(err).Warn("toml read wrong")
	}

	err = models.DBInit(db)
	if err != nil {
		logrus.WithError(err).Warn("DB init wrong")
	}

	err = routers.RouterInit(webport)
	if err != nil {
		logrus.WithError(err).Warn("Router init wrong")
	}
}
