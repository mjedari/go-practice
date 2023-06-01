package main

import (
	"cloud-patterns/retry"
	"cloud-patterns/storage"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

const RETRIY_TIMES = 10
const RETRY_DELAY = time.Second * 2

func main() {

	// want to initiate one instance of storage client
	ctx := context.TODO()

	redisRetry, err := retry.Retry(func(ctx context.Context) (any, error) {

		client, err := storage.NewRedis("localhost", "1212")

		return client, err
	}, RETRIY_TIMES, RETRY_DELAY)(ctx)

	if err != nil {
		panic(err)
	}
	// here we convert interface datatype to redis.Client
	redisRetry.(*redis.Client).Ping(ctx)

}
