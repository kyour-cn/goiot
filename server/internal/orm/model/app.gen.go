// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"encoding/json"
)

const TableNameApp = "app"

// App 应用列表
type App struct {
	ID     int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name   string `gorm:"column:name;not null;comment:应用名称" json:"name"`             // 应用名称
	Key    string `gorm:"column:key;not null;comment:应用KEY 别名" json:"key"`           // 应用KEY 别名
	Remark string `gorm:"column:remark;not null;comment:备注" json:"remark"`           // 备注
	Status int32  `gorm:"column:status;not null;default:1;comment:状态" json:"status"` // 状态
	Sort   int32  `gorm:"column:sort;not null;comment:排序 ASC" json:"sort"`           // 排序 ASC
}

// MarshalBinary 支持json序列化
func (m *App) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// UnmarshalBinary 支持json反序列化
func (m *App) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

// TableName App's table name
func (*App) TableName() string {
	return TableNameApp
}
