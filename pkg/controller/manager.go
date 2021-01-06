package controller

import (
	"net/http"
	"newtest/pkg/models"

	"github.com/sirupsen/logrus"
)

// Manager 测试Manager接口权限
func Manager(w http.ResponseWriter, r *http.Request) {

	id := r.Header.Get("id")

	//根据id获取结构体
	user, err := models.UserQueryByID(id)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	//权限认证
	err = user.ManagerPermission()
	if err != nil {
		w.WriteHeader(403)
		logrus.WithError(err).Info("user permission wrong")
		return
	}

	w.WriteHeader(205)
}
