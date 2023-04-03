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
	stat, err := redisconn.Set(ctx, req.Key, req.Value, exptime).Result()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		return "fail"
	}
	return stat
}
func GetRedis(req GetInterface) string {
	req.Lock.Lock()
	defer req.Lock.Unlock()
	val, err := redisconn.Get(ctx, req.Key).Result()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", req.Key)
		get_err := "key not found"
		return get_err
	}
	return val
}
