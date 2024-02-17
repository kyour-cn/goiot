package test

import (
    "fmt"
    "github.com/go-gourd/gourd/config"
    "gourd/internal/orm/model"
    "gourd/internal/orm/query"
    "gourd/internal/tools"
    "gourd/pkg/cache"
    "testing"
)

func TestRedis(t *testing.T) {

    config.SetConfigDir("../config")

    err := tools.InitDatabase()
    if err != nil {
        t.Error(err)
    }

    userId := int32(1)

    user, err := cache.Remember(fmt.Sprintf("user_info:%d", userId), 60, func() (*model.User, error) {
        return query.User.Where(query.User.ID.Eq(userId)).First()
    })
    if err != nil {
        t.Error(err)
    }

    fmt.Println("用户信息", user)

}
