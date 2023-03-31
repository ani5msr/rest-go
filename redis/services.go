package redis

import (
	"context"
	"log"
	"time"
)

var ctx = context.Background()
var redisconn = RedisConnect()

func SetRedis(key string, value string, exp int64) int64 {
	exptime := time.Duration(exp) * time.Second
	var status int64 = 1
	err := redisconn.Set(ctx, value, key, exptime).Err()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		return 0
	}
	return status
}
func GetRedis(key string) string {
	val, err := redisconn.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		get_err := "key not found"
		return get_err
	}
	return val
}
