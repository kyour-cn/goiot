// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"encoding/json"

	"gorm.io/plugin/soft_delete"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int32                 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username   string                `gorm:"column:username;not null" json:"username"`
	Nickname   string                `gorm:"column:nickname" json:"nickname"`
	Mobile     string                `gorm:"column:mobile" json:"mobile"`
	CreateTime uint                  `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime uint                  `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
	DeleteTime soft_delete.DeletedAt `gorm:"column:delete_time" json:"delete_time"`
}

// MarshalBinary 支持json序列化
func (m *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// UnmarshalBinary 支持json反序列化
func (m *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}