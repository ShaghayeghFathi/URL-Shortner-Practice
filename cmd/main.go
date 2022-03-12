package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/ShaghyeghFathi/URL-Shortner-Practice/internal/db/redis"
	"github.com/ShaghyeghFathi/URL-Shortner-Practice/internal/http"
	"github.com/labstack/echo/v4"
)

func Execute(){
	fmt.Print("hello")
	ctx := context.Background()
	redisClient, err := redis.New(ctx)
	if err != nil {
		log.Printf("error initializing redis client: %v", err)
	}
	rdb := redis.Redis{Client: *redisClient}

	c:= echo.New()
	http.Handler{Redisdb: rdb}.Register(c)
	
	if err := c.Start(":1323"); err != nil {
		log.Fatal(err)
	}
}