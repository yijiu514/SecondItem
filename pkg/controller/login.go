package controller

import (
	"errors"
	"fmt"
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	// ErrPassWordWrong 密码错误
	ErrPassWordWrong = errors.New("the password is wrong")
	// ErrLocked 用户已经被锁定
	ErrLocked = errors.New("the user is locked")
)

// Login 用户登陆功能
func Login(w http.ResponseWriter, r *http.Request) {

	u := GetMassage(r)

	//登陆验证
	err := u.verifylogin()
	if errors.Is(err, encryption.ErrEmailWrong) && errors.Is(err, ErrPassWordWrong) {
		w.WriteHeader(401)
		logrus.WithError(err).Info("the token is invalid")
		return
	} else if errors.Is(err, ErrLocked) {
		w.WriteHeader(423)
		logrus.WithError(err).Info("the user is locked")
		return
	} else if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("serve happen something wrong")
		return
	}

	//获取user信息
	user, err := models.UserQueryByEmail(u.email)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	//令牌下发
	err = encryption.TokenIssue(user, w)
	if err != nil {
		logrus.WithError(err).Warn("token issue wrong in server")
		return
	}

	//放回成功状态码
	w.WriteHeader(201)
}

func (u RegMassage) verifylogin() error {

	//邮箱验证
	if encryption.VerifyEmailFormat(u.email) != nil {
		return encryption.ErrEmailWrong
	}

	//密码验证
	user, err := models.UserQueryByEmail(u.email)
	if err != nil {
		return fmt.Errorf("mysql query wrong %w", err)
	}
	if encryption.Md5Stirng(u.password+user.Passwordsalt) != user.Password {
		return ErrPassWordWrong
	}

	//锁定验证
	if user.Lockat > time.Now().Unix() {
		return ErrLocked
	}

	//登陆成功
	return nil
}
