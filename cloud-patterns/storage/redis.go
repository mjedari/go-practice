package storage

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type Redis struct {
	*redis.Client
}

func NewRedis(host, port string) (*redis.Client, error) {
	log.Print("trying to connect redis on: ", host, port)
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", host, port),
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
