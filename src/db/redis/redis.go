package redis

import (
	"context"
	"fmt"
	"github.com/HuckOps/notify/src/config"
	"github.com/go-redis/redis/v8"
)

type rds struct {
	context    context.Context
	client     *redis.Client
	reloadChan chan bool
}

var Redis *rds

func init() {
	Redis = &rds{}
}

func (r *rds) Load() {
	r.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Config.DB.Redis.Host, config.Config.DB.Redis.Port),
		Password: config.Config.DB.Redis.Password,
		DB:       config.Config.DB.Redis.DB,
	})
	if _, err := r.client.Ping(context.TODO()).Result(); err != nil {
		panic(err)
	}
}

func (r *rds) Client() *redis.Client {
	return r.client
}
