package service

import (
	"encoding/json"
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

// Menu 定义菜单树结构体
type Menu struct {
	PID       int32          `json:"parentId"`
	ID        int32          `json:"id"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	RuleID    int32          `json:"rule_id"`
	Component string         `json:"component"`
	Sort      int32          `json:"sort"`
	Meta      map[string]any `json:"meta"`
	Children  []Menu         `json:"children,omitempty"` // 子菜单列表，omitempty 标签确保为空时不序列化
}

// 递归构建菜单树的函数
func buildMenuTree(menuList []*model.Menu, pid int32) []Menu {
	var result []Menu
	for _, menuM := range menuList {
		if menuM.Pid == pid {

			mate := make(map[string]any)
			if menuM.Meta != "" {
				err := json.Unmarshal([]byte(menuM.Meta), &mate)
				if err != nil {
					log.Error("json解析失败：" + err.Error())
				}
			}

			menu := Menu{
				ID:        menuM.ID,
				PID:       menuM.Pid,
				Name:      menuM.Name,
				Path:      menuM.Path,
				Component: menuM.Component,
				Sort:      menuM.Sort,
				Meta:      mate,
				RuleID:    menuM.RuleID,
			}

			menu.Children = buildMenuTree(menuList, menuM.ID)

			result = append(result, menu)
		}
	}
	return result
}

func GetMenu(appId int32) []Menu {
	m := query.Menu
	menus, err := m.Where(m.AppID.Eq(appId)).Find()
	if err != nil {
		return nil
	}

	list := buildMenuTree(menus, 0)

	return list
}
