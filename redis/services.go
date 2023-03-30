package redis

import (
	"encoding/json"
	"log"
)

func SetRedis(value string, key string, expiry string) int64 {
	redis := RedisConnect() // create the postgres db connection
	var status int64 = 1
	json, _ := json.Marshal(user)
	err := redis.Set(value, key, expiry).Err()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v into Redis", err)
		return 0
	}
	return status
}
