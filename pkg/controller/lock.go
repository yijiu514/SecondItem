package controller

import (
	"errors"
	"net/http"
	"newtest/pkg/encryption"
	"newtest/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// ErrLockSelf 锁定自己
var ErrLockSelf = errors.New("admin can not lock self")

// Lock 锁定用户
func Lock(w http.ResponseWriter, r *http.Request) {

	//获取目标id
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	//token验证
	id, err := TokenCheck(r)
	if errors.Is(err, encryption.ErrTokenWrong) && errors.Is(err, encryption.ErrTokenEmpty) {
		w.WriteHeader(401)
		logrus.WithError(err).Info("somebody do with the token wrong")
		return
	}

	//id判断
	err = judge(id, userID)
	if err != nil {
		w.WriteHeader(403)
		logrus.WithError(err).Info("the user want lock himself")
		return
	}

	//锁定user
	user, err := models.UserQueryByID(userID)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}
	user.UpdateLockat(99999999)
}

// UnLock 解锁用户
func UnLock(w http.ResponseWriter, r *http.Request) {
	//获取目标id
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	//token验证
	_, err := TokenCheck(r)
	if errors.Is(err, encryption.ErrTokenWrong) && errors.Is(err, encryption.ErrTokenEmpty) {
		w.WriteHeader(401)
		logrus.WithError(err).Info("somebody do with the token wrong")
		return
	}

	//锁定user
	user, err := models.UserQueryByID(userID)
	if err != nil {
		w.WriteHeader(500)
		logrus.WithError(err).Warn("mysql query wrong")
		return
	}

	user.UpdateLockat(0)
}

//判断id是否相同
func judge(id int, idself int) (err error) {
	if id == idself {
		return ErrLockSelf
	}
	return nil
}
