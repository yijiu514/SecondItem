package controller

import (
	"errors"
	"fmt"
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"
	"time"

	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
)

var decoder = schema.NewDecoder()

// RegMassage 获取信息并转换成结构体
type RegMassage struct {
	email    string
	password string
}

// Register 用户注册
func Register(w http.ResponseWriter, r *http.Request) {

	//获取注册信息
	u := GetMassage(r)

	//注册验证
	err := verifyregister(u.email)
	if errors.Is(err, encryption.ErrEmailWrong) {
		w.WriteHeader(401)
		logrus.WithError(err).Info("the user email wrong")
		return
	} else if errors.Is(err, models.ErrUserISExist) {
		w.WriteHeader(409)
		logrus.WithError(err).Info("the user email is exist")
		return
	} else if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("somothing wrong happend in server")
		return
	}

	//初始化user结构体
	user, err := u.regWrite()
	if err != nil {
		logrus.WithError(err).Warn("mysql query wrong in server")
		return
	}

	//写入数据库
	err = user.Insert()
	if err != nil {
		logrus.WithError(err).Warn("mysql insert wrong in server")
		return
	}

	//下发令牌
	err = encryption.TokenIssue(user, w)
	if err != nil {
		logrus.WithError(err).Warn("token issue wrong in server")
		return
	}

}

// GetMassage 获取前端信息
func GetMassage(r *http.Request) (msg RegMassage) {
	r.ParseForm()
	err := decoder.Decode(&msg, r.PostForm)
	if err != nil {
		logrus.WithError(err).Warn("shema get msg wrong")
		return RegMassage{}
	}
	return
}

//注册验证
func verifyregister(email string) (err error) {

	//邮箱验证
	if encryption.VerifyEmailFormat(email) != nil {
		return fmt.Errorf("register wrong %w", encryption.ErrEmailWrong)
	}

	//数据库验证
	if models.IsUserExist(email) != nil {
		return fmt.Errorf("register wrong %w", models.ErrUserISExist)
	}

	return nil
}

//初始化信息结构体
func (u RegMassage) regWrite() (user models.User, err error) {

	user.Email = u.email
	user.Password, user.Passwordsalt = encryption.Md5Salt(u.password, 8)
	user.Creatat = time.Now().Unix()
	user.Lockat = 0
	maxid, err := models.QueryMaxID()
	user.ID = maxid + 1
	user.Sessionsalt = encryption.RandomString(8)
	user.Role = "editor"

	return
}
