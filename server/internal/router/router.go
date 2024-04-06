package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-gourd/gourd/config"
	adminRoute "gourd/internal/app/admin/route"
	"gourd/internal/router/middleware"
	"net/http"
	"os"
)

var router *chi.Mux

func GetRouter() *chi.Mux {

	if router != nil {
		return router
	}

	router = chi.NewRouter()

	return router
}

// Register 注册路由
func Register() {

	r := GetRouter()

	// 跨域
	r.Use(middleware.CORS)

	// 主页
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello gourd!"))
	})

	// 404响应
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {

		// 若路由未定义，检测是否为静态资源
		conf := config.GetHttpConfig()
		if conf.Static != "" {
			filepath := conf.Static + r.URL.Path
			//判断文件是否存在
			_, err := os.Stat(filepath)
			if err == nil {
				http.ServeFile(w, r, filepath)
				return
			}
		}

		// 404响应内容
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404 not found."))
	})

	// 注册admin相关路由
	adminGroup := chi.NewRouter().
		Group(adminRoute.RegisterRoute)
	r.Mount("/admin", adminGroup)

}
