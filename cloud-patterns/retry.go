package main

import (
	"cloud-patterns/retry"
	"cloud-patterns/storage"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

const RetryTimes = 10
const RetryDelay = time.Second * 2

func main() {

	// want to initiate one instance of storage client
	ctx := context.TODO()

	redisRetry, err := retry.Retry(func(ctx context.Context) (any, error) {

		client, err := storage.NewRedis("localhost", "1212")

		return client, err
	}, RetryTimes, RetryDelay)(ctx)

	if err != nil {
		panic(err)
	}
	// here we convert interface datatype to redis.Client
	redisRetry.(*redis.Client).Ping(ctx)

}
