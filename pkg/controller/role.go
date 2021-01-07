package controller

import (
	"net/http"
	"newtest/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
)

// Role 修改角色信息
func Role(w http.ResponseWriter, r *http.Request) {

	//获取目标信息
	userID := (chi.URLParam(r, "userID"))
	role := r.Form.Get("role")
	id, _ := strconv.Atoi(userID)

	//写入数据库
	user, err := models.UserQueryByID(id)
	AssertErr(500, err)

	//更新数据库信息
	err = user.UpdateRole(role)
	AssertErr(500, err)
}
