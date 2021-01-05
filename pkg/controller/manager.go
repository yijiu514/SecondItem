package controller

import (
	"errors"
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"

	"github.com/sirupsen/logrus"
)

// Manager 测试Manager接口权限
func Manager(w http.ResponseWriter, r *http.Request) {

	//token认证
	id, err := TokenCheck(r)
	if errors.Is(err, encryption.ErrTokenWrong) && errors.Is(err, encryption.ErrTokenEmpty) {
		w.WriteHeader(401)
		logrus.WithError(err).Info("somebody do with the token wrong")
		return
	}

	//根据id获取结构体
	user, err := models.UserQueryByID(id)
	if err != nil {
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	//权限认证
	err = user.ManagerPermission()
	if err != nil {
		logrus.WithError(err).Info("user permission wrong")
		return
	}

	w.WriteHeader(205)
}
