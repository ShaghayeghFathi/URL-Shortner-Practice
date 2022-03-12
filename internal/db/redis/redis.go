package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
)

type Redis struct{
	Client redis.Client
}

func New(ctx context.Context)(*redis.Client,error){
	timeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	log.Info("Connecting to redis...")
	if err := client.Ping(timeout).Err(); err != nil {
		return nil, fmt.Errorf("error connecting to redis: %v", err)
	}
	return client,nil
}

func (r Redis) Set(ctx context.Context,key string, value string)error{
	return r.Client.Set(ctx, key, value, 0).Err()
}