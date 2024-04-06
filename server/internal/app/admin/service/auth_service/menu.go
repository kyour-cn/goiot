package auth_service

import (
	"encoding/json"
	"github.com/go-gourd/gourd/log"
	"gourd/internal/orm/model"
	"gourd/internal/orm/query"
)

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
