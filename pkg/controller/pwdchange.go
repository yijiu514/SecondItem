package controller

import (
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"
)

// PasswordChange 密码修改
func PasswordChange(w http.ResponseWriter, r *http.Request) {

	//获取id和token信息

	pwd := r.Form.Get("newpassword")

	//根据id获取信息结构体
	id := r.Context().Value("id").(int)
	user, err := models.UserQueryByID(id)
	AssertErr(500, err)

	//生成新密码
	newpwd, salt := encryption.Md5Salt(pwd, 8)

	//写入数据库
	err = user.UpdatePwdChange(newpwd, salt)
	AssertErr(500, err)

	w.WriteHeader(204)
}
