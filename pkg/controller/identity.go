package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newtest/pkg/models"
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
	AssertErr(500, err)

	//生成梵返回信息
	msg, err := identitymessage(user)
	AssertErr(500, err)

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
		fmt.Errorf("get json wrong %m", err)
	}
	json := string(jsonbyte)
	return json, nil
}
