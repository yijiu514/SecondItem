package controller

import (
	"errors"
	"net/http"
	"newtest/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
)

// ErrLockSelf 锁定自己
var ErrLockSelf = errors.New("admin can not lock self")

// Lock 锁定用户
func Lock(w http.ResponseWriter, r *http.Request) {

	//获取目标id
	userID := chi.URLParam(r, "userID")

	//id判断
	id := r.Context().Value("id").(int)
	userid, _ := strconv.Atoi(userID)
	err := judge(id, userid)
	AssertErr(403, err)

	//数据库查询
	user, err := models.UserQueryByID(userid)
	AssertErr(500, err)

	//锁定user
	err = user.UpdateLockat(99999999)
	AssertErr(500, err)

	w.WriteHeader(201)
}

// UnLock 解锁用户
func UnLock(w http.ResponseWriter, r *http.Request) {

	//获取目标id
	userID := chi.URLParam(r, "userID")

	//锁定user
	id, _ := strconv.Atoi(userID)
	user, err := models.UserQueryByID(id)
	AssertErr(500, err)

	err = user.UpdateLockat(0)
	AssertErr(500, err)
}

//判断id是否相同
func judge(id int, idself int) (err error) {

	if id == idself {
		return ErrLockSelf
	}
	return nil
}
