package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newtest/pkg/models"

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

	id := r.Context().Value("id").(int)

	//根据id查询信息结构体
	user, err := models.UserQueryByID(id)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	//生成梵返回信息
	msg, err := identitymessage(user)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("message creat wrong")
		return
	}

	fmt.Fprintln(w, msg)
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
