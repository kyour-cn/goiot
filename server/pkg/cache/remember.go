package cache

import (
	"context"
	"encoding/json"
	"gourd/internal/tools"
	"time"
)

// Remember 从缓存中获取数据，如果不存在则从数据库中获取数
// 数据回调函数的返回值必须是支持json序列化的结构体
func Remember[T any](key string, ttl int, dataCallback func() (*T, error)) (*T, error) {

	redis, err := tools.GetRedis()
	if err != nil {
		return nil, err
	}

	var ctx = context.Background()
	// Try to get the value from the cache
	value, err := redis.Get(ctx, key).Result()
	if err == nil {
		// If the value exists in the cache, deserialize and return it
		var result T
		err = json.Unmarshal([]byte(value), &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}

	// If the value doesn't exist in the cache, call the getDataFunc to fetch the value
	value2, err := dataCallback()
	if err != nil {
		return nil, err
	}

	// Serialize the value
	serializedValue, err := json.Marshal(value2)
	if err != nil {
		return nil, err
	}

	// Store the serialized value in the cache with the specified TTL
	err = redis.Set(ctx, key, serializedValue, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		return nil, err
	}

	return value2, nil
}
