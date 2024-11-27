package redis

import (
	"context"
	"fmt"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s%d", config.Global.Redis.Host, config.Global.Redis.Port), // Redis server address
		Password: "",                                                                      // no password set
		DB:       0,                                                                       // default DB
	})
	er := rdb.Ping(context.Background()).Err()
	if er != nil {
		panic(er)
	}
	return rdb
}
