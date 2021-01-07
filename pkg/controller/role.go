package controller

import (
	"net/http"
	"newtest/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// Role 修改角色信息
func Role(w http.ResponseWriter, r *http.Request) {

	//获取目标信息
	userID := (chi.URLParam(r, "userID"))
	role := r.Form.Get("role")
	id, _ := strconv.Atoi(userID)
	//写入数据库
	user, err := models.UserQueryByID(id)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	//更新数据库信息
	err = user.UpdateRole(role)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql update wrong")
		return
	}
}
