package redis

import (
	"context"
	"log"
	"time"
)

var ctx = context.Background()
var redisconn = RedisConnect()

func SetRedis(key string, value string, exp int64) string {
	exptime := time.Duration(exp) * time.Second
	stat, err := redisconn.Set(ctx, key, value, exptime).Result()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		return "fail"
	}
	return stat
}
func GetRedis(key string) string {
	val, err := redisconn.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", key)
		get_err := "key not found"
		return get_err
	}
	return val
}
