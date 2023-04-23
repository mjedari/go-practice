package main

import (
	"cloud-patterns/retry"
	"cloud-patterns/storage"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {

	// want to initiate one instance of storage client
	ctx := context.TODO()

	redisRetry, err := retry.Retry(func(ctx context.Context) (any, error) {

		client, err := storage.NewRedis("localhost", "1212")

		return client, err
	}, 10, time.Second*2)(ctx)

	if err != nil {
		panic(err)
	}
	redisRetry.(*redis.Client).Ping(ctx)

}
