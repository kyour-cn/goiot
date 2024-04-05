package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"gourd/internal/orm/model"
	"gourd/internal/orm/query"
	"time"
)

// LoginService 登录服务
type LoginService struct {
}

// GenToken 生成token
func (us *LoginService) GenToken(username string, password string) (*model.User, string, error) {

	u := query.User

	// 查询符合条件的用户
	user, err := u.Where(
		u.Username.Eq(username),
		u.Password.Eq(password),
	).First()
	if err != nil {
		return nil, "", errors.New("用户名或密码错误")
	}

	var mySigningKey = []byte("your_secret_key")

	now := time.Now()
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)), // Token expires after 24 hours
		Issuer:    "your_issuer",
		Subject:   string(user.ID), // Assuming userID is the subject of the token
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenVal, err := token.SignedString(mySigningKey)
	if err != nil {
		return nil, "", err
	}

	// 生成token
	return user, tokenVal, nil
}
