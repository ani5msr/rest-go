package redis

import (
	"context"
	"log"
	"sync"
	"time"
)

var ctx = context.Background()
var redisconn = RedisConnect()
var mu sync.RWMutex

func SetRedis(req SetInterface) string {
	exptime := time.Duration(req.Expiry) * time.Second
	req.Lock.Lock()
	defer req.Lock.Unlock()
	defer mu.Unlock()
	stat, err := redisconn.Set(ctx, req.Key, req.Value, exptime).Result()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		return "fail"
	}
	return stat
}
func GetRedis(key string) string {
	mu.Lock()
	defer mu.Unlock()
	val, err := redisconn.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", key)
		get_err := "key not found"
		return get_err
	}
	return val
}
