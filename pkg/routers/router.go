package routers

import (
	"net/http"
	"newtest/pkg"
	"newtest/pkg/controller"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func init() {
	//设置日志输出格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	//自定义路由
	r := chi.NewRouter()
	r.Post("/login", controller.Login)
	r.Post("/register", controller.Register)
	r.Post("/my/password", controller.PasswordChange)
	r.Get("/identity", controller.Identity)

	r.Route("/{userID}", func(r chi.Router) {
		r.Post("/lock", controller.Lock)
		r.Delete("/lock", controller.UnLock)
		r.Put("/role", controller.Role)
		r.Delete("/password", controller.PasswordSet)
	})

	r.Get("/test/editor", controller.Editor)
	r.Get("/test/manager", controller.Manager)

	//从配置文件获取webport信息
	w, _, err := pkg.TomlRead()
	if err != nil {
		logrus.WithError(err).Warn("toml read wrong")
	}
	webport := strconv.Itoa(int(w))

	//监听
	err = http.ListenAndServe(":"+webport, r)
	if err != nil {
		logrus.WithError(err).Warn("http listen wrong")
	}
}
