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
		AssertErr(401, err)
	} else if errors.Is(err, models.ErrUserISExist) {
		AssertErr(409, err)
	} else if err != nil {
		AssertErr(500, err)
	}

	//初始化user结构体
	user, err := u.regWrite()
	AssertErr(500, err)

	//写入数据库
	err = user.Insert()
	AssertErr(500, err)

	//下发令牌
	err = encryption.TokenIssue(user, w)
	AssertErr(500, err)

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
	user.Sessionsalt, _ = encryption.RandomString(8)
	user.Role = "editor"

	return
}
