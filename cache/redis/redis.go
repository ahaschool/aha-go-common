package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type Config struct {
	Host        string
	Auth        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Db          int
}

// NewRedis
func NewRedis(config Config) *redis.Pool {
	option := redis.DialPassword(config.Auth)
	RedisPool := &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: time.Duration(config.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			client, err := redis.Dial("tcp", config.Host, option)
			if err != nil {
				return client, err
			}
			_ ,err = client.Do("SELECT", config.Db)
			return client, err
		},
	}
	return RedisPool
}
