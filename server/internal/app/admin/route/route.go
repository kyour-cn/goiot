package route

import (
	"github.com/go-chi/chi/v5"
	loginCtl "gourd/internal/app/admin/ctl"
)

// RegisterRoute 注册路由组
func RegisterRoute(r chi.Router) {

	lCtl := loginCtl.AuthCtl{}
	r.HandleFunc("/auth/captcha", lCtl.Captcha)
	r.HandleFunc("/auth/login", lCtl.Login)
	r.HandleFunc("/auth/menu", lCtl.Menu)

}
