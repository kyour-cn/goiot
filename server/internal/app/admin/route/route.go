package route

import (
	"github.com/go-chi/chi/v5"
	loginCtl "gourd/internal/app/admin/ctl"
	"gourd/internal/app/admin/ctl/system"
)

// RegisterRoute 注册路由组
func RegisterRoute(r chi.Router) {

	lCtl := loginCtl.AuthCtl{}
	r.HandleFunc("/auth/captcha", lCtl.Captcha)
	r.HandleFunc("/auth/login", lCtl.Login)
	r.HandleFunc("/auth/menu", lCtl.Menu)

	// 注册admin相关路由
	adminGroup := chi.NewRouter().
		Group(system.RegisterRoute)
	r.Mount("/system", adminGroup)

}
