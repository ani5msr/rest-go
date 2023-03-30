package redis

import (
	"log"
	"time"
)

func SetRedis(key string, value string, exp int64) int64 {

	exptime := time.Duration(exp) * time.Second
	redis := RedisConnect()
	var status int64 = 1
	err := redis.Set(value, key, exptime).Err()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		return 0
	}
	return status
}
func GetRedis(key string) string {
	redis := RedisConnect()
	val, err := redis.Get(key).Result()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		get_err := "key not found"
		return get_err
	}
	return val
}
