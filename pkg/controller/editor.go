package controller

import (
	"net/http"
	"newtest/pkg/models"

	"github.com/sirupsen/logrus"
)

// Editor 测试editor接口权限
func Editor(w http.ResponseWriter, r *http.Request) {

	//id获取
	id := r.Header.Get("id")

	//根据id获取结构体
	user, err := models.UserQueryByID(id)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	//权限认证
	err = user.EditorPermission()
	if err != nil {
		w.WriteHeader(403)
		logrus.WithError(err).Info("user permission wrong")
		return
	}

	w.WriteHeader(205)
}
