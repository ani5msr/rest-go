package redis

import (
	"log"
)

func SetRedis(words string) int64 {
	redis := RedisConnect()
	var status int64 = 1
	err := redis.Set(value, key, expiry).Err()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		return 0
	}
	return status
}
