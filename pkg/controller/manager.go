package controller

import (
	"net/http"
	"newtest/pkg/models"
)

// Manager 测试Manager接口权限
func Manager(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("id").(int)

	//根据id获取结构体
	user, err := models.UserQueryByID(id)
	AssertErr(500, err)

	//权限认证
	err = user.ManagerPermission()
	AssertErr(403, err)

	w.WriteHeader(205)
}
