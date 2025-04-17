package connections

import (
	"context"
	"github.com/manicar2093/gomancer/bootstrap/core/httphealthcheck"
	"log"

	"github.com/redis/go-redis/v9"
)

type (
	RedisConfig struct {
		Dns string `env:"REDIS_URL" validate:"required|isURL"`
	}

	RedisWrapper struct {
		*redis.Client
	}
)

func GetRedis(config RedisConfig) *RedisWrapper {
	opts, err := redis.ParseURL(config.Dns)
	if err != nil {
		log.Panic(err)
	}

	client := redis.NewClient(opts)

	return &RedisWrapper{
		Client: client,
	}
}

func (c *RedisWrapper) ServiceHealth() (httphealthcheck.HealthStatusData, error) {
	ping := c.Client.Ping(context.Background())
	if ping.Err() != nil {
		return httphealthcheck.HealthStatusData{
			Error: ping.Err(),
		}, nil
	}

	return httphealthcheck.HealthStatusData{
		IsAvailable: true,
	}, nil
}

func (c *RedisWrapper) ServiceName() string {
	return "redis_connection"
}
