package utilities

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetRedisKey(rdb *redis.Client, ctx context.Context, key string, value any, ttl time.Duration) {
	go func() {
		data, err := json.Marshal(value)
		if err != nil {
			log.Println("error marshalling value", err)
		}

		err = rdb.Set(ctx, key, data, ttl).Err()
		if err != nil {
			log.Println("redis error setting key", err)
		}
	}()
}

func GetUnmarshalRedisKey(rdb *redis.Client, ctx context.Context, key string, dst any) error {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), dst)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRedisKey(rdb *redis.Client, ctx context.Context, key string) error {
	return rdb.Del(ctx, key).Err()
}
