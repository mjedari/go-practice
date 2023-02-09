package infra

import "github.com/redis/go-redis/v9"

type RedisStorage struct {
	redis *redis.Client
}

func (r RedisStorage) Persist(bytes []byte) error {
	// store data in redis in string format or byte?!!
	// for log its best to store as a set of strings
	return nil
}
