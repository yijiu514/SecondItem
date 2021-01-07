package controller

import (
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
)

// PasswordSet 重置密码
func PasswordSet(w http.ResponseWriter, r *http.Request) {

	//获取目标id
	userID := chi.URLParam(r, "userID")

	//写入数据库
	id, _ := strconv.Atoi(userID)
	user, err := models.UserQueryByID(id)
	AssertErr(500, err)

	err = user.UpdatePwdChange(encryption.Md5Salt("123456", 8))
	AssertErr(500, err)
}
