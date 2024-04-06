package auth_service

import (
	"errors"
	"fmt"
	"github.com/go-gourd/gourd/config"
	"github.com/go-gourd/gourd/log"
	"github.com/golang-jwt/jwt/v5"
	"gourd/internal/orm/model"
	"gourd/internal/orm/query"
	"strconv"
	"time"
)

type JwtConfig struct {
	Secret string `toml:"secret" json:"secret"` // 密钥
	Expire uint   `toml:"expire" json:"expire"` // 过期时间 / 秒
}

var jwtConfig *JwtConfig

func init() {
	jwtConfig = &JwtConfig{}

	// 读取配置
	err := config.Unmarshal("jwt", jwtConfig)
	if err != nil {
		log.Warn("jwt配置读取失败：" + err.Error())
	}
}

// GenToken 生成token
func GenToken(username string, password string) (*model.User, string, error) {

	u := query.User

	// 查询符合条件的用户
	user, err := u.Where(
		u.Username.Eq(username),
		u.Password.Eq(password),
	).First()
	if err != nil {
		return nil, "", errors.New("用户名或密码错误")
	}

	now := time.Now()
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Second * time.Duration(jwtConfig.Expire))),
		Subject:   strconv.Itoa(int(user.ID)),
		ID:        strconv.Itoa(int(user.ID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenVal, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return nil, "", err
	}

	// 生成token
	return user, tokenVal, nil
}

// ParseToken 解析token
func ParseToken(token string) (*jwt.MapClaims, error) {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtConfig.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok {
		return &claims, nil
	} else {
		return nil, errors.New("token解析失败")
	}
}

// CheckAuth 校验token并返回用户ID
func CheckAuth(auth string) (uid int32, err error) {

	// 校验token前缀是否`Bearer `
	if len(auth) < 7 || auth[:7] != "Bearer " {
		err = errors.New("token格式错误")
		return
	}

	// 去除token前缀
	auth = auth[7:]
	if token, err := ParseToken(auth); err != nil {
		return 0, err
	} else {
		uidS, _ := token.GetSubject()
		uid, err := strconv.Atoi(uidS)
		return (int32)(uid), err
	}
}
