package ctl

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/steambap/captcha"
	"gourd/internal/app/admin"
	"gourd/internal/app/admin/service"
	"gourd/pkg/cache"
	"net/http"
)

// AuthCtl 测试
type AuthCtl struct {
	admin.BaseController //继承基础控制器
}

// Captcha 测试
func (ctl *AuthCtl) Captcha(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	// 获取number参数
	number := params.Get("number")
	if number == "" {
		_ = ctl.Fail(w, http.StatusBadRequest, "number is empty", nil)
		return
	}

	img, err := captcha.New(160, 50, func(options *captcha.Options) {
		options.CharPreset = "1234567890"
	})
	if err != nil {
		_ = ctl.Fail(w, http.StatusInternalServerError, "new captcha error", err.Error())
		return
	}

	// 缓存验证码
	err = cache.Set("admin:login_code:"+number, img.Text, 60)
	if err != nil {
		_ = ctl.Fail(w, http.StatusInternalServerError, "set cache error", err.Error())
		return
	}

	_ = img.WriteImage(w)

}

// Login 测试
func (ctl *AuthCtl) Login(w http.ResponseWriter, r *http.Request) {

	number := r.FormValue("number")
	code := r.FormValue("code")
	if number == "" || code == "" {
		_ = ctl.Fail(w, http.StatusBadRequest, "验证码未填写", nil)
		return
	}

	codeCacheKey := "admin:login_code:" + number
	if cacheCode, err := cache.Get(codeCacheKey); err != nil || cacheCode != code {
		_ = cache.Del(codeCacheKey)
		_ = ctl.Fail(w, http.StatusBadRequest, "验证码错误", nil)
		return
	}
	_ = cache.Del(codeCacheKey)

	// 登录
	username := r.FormValue("username")
	password := r.FormValue("password")

	// 判断密码是否md5
	if r.FormValue("md5") != "true" {
		// 密码需要转md5
		pwdMd5 := md5.New()
		pwdMd5.Write([]byte(password))
		password = hex.EncodeToString(pwdMd5.Sum(nil))
	}

	if user, token, err := service.GenToken(username, password); err != nil {
		_ = ctl.Fail(w, http.StatusInternalServerError, err.Error(), nil)
		return
	} else {

		data := map[string]any{
			"token": token,
			"userInfo": map[string]any{
				"id":       user.ID,
				"username": user.Username,
				"nickname": user.Nickname,
				"mobile":   user.Mobile,
				"avatar":   user.Avatar,
			},
		}

		_ = ctl.Success(w, "", data)
		return
	}
}

// Menu 获取菜单
func (ctl *AuthCtl) Menu(w http.ResponseWriter, r *http.Request) {

	auth := r.Header.Get("Authorization")

	// 校验token前缀是否`Bearer `
	if len(auth) < 7 || auth[:7] != "Bearer " {
		_ = ctl.Fail(w, http.StatusBadRequest, "token格式错误", nil)
		return
	}

	// 去除token前缀
	auth = auth[7:]
	if token, err := service.ParseToken(auth); err != nil {
		_ = ctl.Fail(w, http.StatusInternalServerError, err.Error(), nil)
	} else {
		_ = token
		menu := service.GetMenu(1)
		data := map[string]any{
			"menu":        menu,
			"permissions": []string{},
		}
		_ = ctl.Success(w, "", data)
	}

}
