package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Ident 个人信息返回
type Ident struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Creatat int64  `json:"creatat"`
	Role    string `json:"role"`
}

// Identity 返回个人信息功能
func Identity(w http.ResponseWriter, r *http.Request) {

	id, err := TokenCheck(r)
	if errors.Is(err, encryption.ErrTokenWrong) && errors.Is(err, encryption.ErrTokenEmpty) {
		w.WriteHeader(401)
		logrus.WithError(err).Info("somebody do with the token wrong")
		return
	}

	//根据id查询信息结构体
	user, err := models.UserQueryByID(id)
	if err != nil {
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}
	//生成梵返回信息
	msg, err := identitymessage(user)
	if err != nil {
		logrus.WithError(err).Warn("message creat wrong")
		return
	}

	fmt.Fprintln(w, msg)

}

// TokenCheck 实现token的认证
func TokenCheck(r *http.Request) (id int, err error) {
	r.ParseForm()
	head := r.Header
	tokenStr := head.Get("token")
	id, _ = strconv.Atoi(head.Get("id"))
	err = encryption.TokenVerify(tokenStr, id)
	return
}

func identitymessage(user models.User) (message string, err error) {
	var u Ident
	u.ID = user.ID
	u.Email = user.Email
	u.Creatat = user.Creatat
	u.Role = user.Role
	jsonbyte, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		fmt.Println("getjson wrong")
	}
	json := string(jsonbyte)
	return json, nil
}
