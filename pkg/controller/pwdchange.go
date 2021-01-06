package controller

import (
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"

	"github.com/sirupsen/logrus"
)

// PasswordChange 密码修改
func PasswordChange(w http.ResponseWriter, r *http.Request) {

	//获取id和token信息
	id := r.Header.Get("id")
	pwd := r.Header.Get("newpassword")

	//根据id获取信息结构体
	user, err := models.UserQueryByID(id)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	//生成新密码
	newpwd, salt := encryption.Md5Salt(pwd, 8)

	//写入数据库
	err = user.UpdatePwdChange(newpwd, salt)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql update wrong")
		return
	}

	w.WriteHeader(204)
}
