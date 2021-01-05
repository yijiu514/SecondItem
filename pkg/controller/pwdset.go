package controller

import (
	"errors"
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// PasswordSet 重置密码
func PasswordSet(w http.ResponseWriter, r *http.Request) {
	//获取目标id
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	//获取角色信息
	r.ParseForm()

	//token验证
	_, err := TokenCheck(r)
	if errors.Is(err, encryption.ErrTokenWrong) && errors.Is(err, encryption.ErrTokenEmpty) {
		w.WriteHeader(401)
		logrus.WithError(err).Info("somebody do with the token wrong")
		return
	}

	//写入数据库
	user, err := models.UserQueryByID(userID)
	if err != nil {
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	user.UpdatePwdChange(encryption.Md5Salt("123456", 8))

	return
}
