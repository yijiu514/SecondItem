package controller

import (
	"errors"
	"net/http"
	"newtest/pkg/models"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// ErrLockSelf 锁定自己
var ErrLockSelf = errors.New("admin can not lock self")

// Lock 锁定用户
func Lock(w http.ResponseWriter, r *http.Request) {

	//获取目标id
	userID := chi.URLParam(r, "userID")

	//id判断
	id := r.Header.Get("id")
	err := judge(id, userID)
	if err != nil {
		w.WriteHeader(403)
		logrus.WithError(err).Info("the user want lock himself")
		return
	}

	user, err := models.UserQueryByID(userID)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	//锁定user
	err = user.UpdateLockat(99999999)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql update  wrong")
	}
}

// UnLock 解锁用户
func UnLock(w http.ResponseWriter, r *http.Request) {

	//获取目标id
	userID := chi.URLParam(r, "userID")

	//锁定user
	user, err := models.UserQueryByID(userID)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	err = user.UpdateLockat(0)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql update  wrong")
	}
}

//判断id是否相同
func judge(id string, idself string) (err error) {

	if id == idself {
		return ErrLockSelf
	}
	return nil
}
