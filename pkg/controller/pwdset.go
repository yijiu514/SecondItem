package controller

import (
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// PasswordSet 重置密码
func PasswordSet(w http.ResponseWriter, r *http.Request) {

	//获取目标id
	userID := chi.URLParam(r, "userID")

	//写入数据库
	user, err := models.UserQueryByID(userID)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	err = user.UpdatePwdChange(encryption.Md5Salt("123456", 8))
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql update wrong")
	}
}
