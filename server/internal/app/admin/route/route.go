package route

import (
	"github.com/go-chi/chi/v5"
	loginCtl "gourd/internal/app/admin/ctl"
)

// RegisterRoute 注册路由组
func RegisterRoute(r chi.Router) {

	lCtl := loginCtl.LoginCtl{}
	r.HandleFunc("/login/captcha", lCtl.Captcha)
	r.HandleFunc("/login/login", lCtl.Login)

}
