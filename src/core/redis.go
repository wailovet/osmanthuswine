package core

import (
	"github.com/go-redis/redis"
)

var instanceRedis *redis.Client

func GetRedis() *redis.Client {
	if instanceRedis == nil {
		config := GetInstanceConfig()
		instanceRedis = redis.NewClient(&redis.Options{
			Addr:     config.Redis.Addr,
			Password: config.Redis.Password, // no password set
			DB:       config.Redis.Db,       // use default DB
		})
	}
	return instanceRedis
}

func init() {
	println("init in redis ")
}