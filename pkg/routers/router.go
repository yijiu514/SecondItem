package routers

import (
	"errors"
	"fmt"
	"net/http"
	"newtest/pkg/controller"
	"newtest/pkg/encryption"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// RouterInit 路由初始化
func RouterInit(w int64) (err error) {

	//设置日志输出格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	//自定义路由
	r := chi.NewRouter()

	//登陆注册
	r.Post("/login", controller.Login)
	r.Post("/register", controller.Register)

	//业务功能
	r.Route("/", func(r chi.Router) {
		r.Use(MiddleTokenCheck)
		r.Post("/my/password", controller.PasswordChange)
		r.Get("/identity", controller.Identity)
		r.Get("/test/editor", controller.Editor)
		r.Get("/test/manager", controller.Manager)
	})

	//管理员权限
	r.Route("/{userID}", func(r chi.Router) {
		r.Use(MiddleTokenCheck)
		r.Post("/lock", controller.Lock)
		r.Delete("/lock", controller.UnLock)
		r.Put("/role", controller.Role)
		r.Delete("/password", controller.PasswordSet)
	})

	//获取webport信息
	webport := strconv.Itoa(int(w))

	//监听
	err = http.ListenAndServe(":"+webport, r)
	return fmt.Errorf("listen wrong %x", err)
}

// MiddleTokenCheck 令牌确认中间件
func MiddleTokenCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := controller.TokenCheck(r)
		if errors.Is(err, encryption.ErrTokenWrong) && errors.Is(err, encryption.ErrTokenEmpty) {
			w.WriteHeader(401)
			return
		}
		next.ServeHTTP(w, r)
	})
}
