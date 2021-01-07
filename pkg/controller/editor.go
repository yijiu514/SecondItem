package controller

import (
	"net/http"
	"newtest/pkg/models"
)

// MyError 定义预Panic结构体
type MyError struct {
	Code int
	Err  error
}

// Editor 测试editor接口权限
func Editor(w http.ResponseWriter, r *http.Request) {

	//id获取
	id := r.Context().Value("id").(int)

	//根据id获取结构体
	user, err := models.UserQueryByID(id)
	AssertErr(500, err)

	//权限认证
	err = user.EditorPermission()
	AssertErr(403, err)

	w.WriteHeader(205)
}

// AssertErr 将错误panic出去
func AssertErr(code int, err error) {
	var res MyError
	res.Code = code
	res.Err = err
	panic(res)
}
