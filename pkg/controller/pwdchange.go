package controller

import (
	"errors"
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"
	"strconv"

	"github.com/sirupsen/logrus"
)

// PasswordChange 密码修改
func PasswordChange(w http.ResponseWriter, r *http.Request) {

	//获取id和token信息
	id, pwd, err := TokenCheckForChange(r)
	if errors.Is(err, encryption.ErrTokenWrong) && errors.Is(err, encryption.ErrTokenEmpty) {
		w.WriteHeader(401)
		logrus.WithError(err).Info("somebody do with the token wrong")
		return
	}

	//根据id获取信息结构体
	user, err := models.UserQueryByID(id)
	if err != nil {
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	//生成新密码
	newpwd, salt := encryption.Md5Salt(pwd, 8)

	//写入数据库
	user.UpdatePwdChange(newpwd, salt)

	w.WriteHeader(204)
}

// TokenCheckForChange 获取token，密码等信息
func TokenCheckForChange(r *http.Request) (id int, newpwd string, err error) {
	r.ParseForm()
	head := r.Header
	tokenStr := head.Get("token")
	id, _ = strconv.Atoi(head.Get("id"))
	err = encryption.TokenVerify(tokenStr, id)
	newpwd = r.Form.Get("newpassword")
	return
}
