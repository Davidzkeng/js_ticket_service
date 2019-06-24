package database

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"js_ticket_service/config"
	"os"
	"time"
)

type redisConfig struct {
	Addr        string
	Auth        string
	Database    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisClient *redis.Pool

func readRedisConfig() *redisConfig {
	redisConfig := redisConfig{
		Addr:        config.Cfg.GetString("redis.addr"),
		Auth:        config.Cfg.GetString("redis.auth"),
		Database:    config.Cfg.GetString("redis.db"),
		MaxIdle:     config.Cfg.GetInt("redis.max_idle"),
		MaxActive:   config.Cfg.GetInt("redis.max_active"),
		IdleTimeout: 180 * time.Second,
	}
	return &redisConfig
}

func OpenRedis() {
	redisConfig := readRedisConfig()
	RedisClient = &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,
		MaxActive:   redisConfig.MaxActive,
		IdleTimeout: redisConfig.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConfig.Addr)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				os.Exit(1)
			}
			if redisConfig.Auth != "" {
				if _, err := c.Do("AUTH", redisConfig.Auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			c.Do("SELECT", redisConfig.Database)
			return c, nil
		},
	}
}

func GetRedis() *redis.Pool {
	OpenRedis()
	return RedisClient
}

func Zadd(key string, score int, member string) error {
	r := GetRedis().Get()
	defer r.Close()
	_, err := r.Do("zAdd", key, score, member)
	if err != nil {
		return err
	}
	return nil
}

func HGet(key, field string) (string, error) {
	r := GetRedis().Get()
	defer r.Close()
	res, err := redis.String(r.Do("hget", key, field))
	if err != nil {
		return "", err
	}
	return res, nil
}

func Get(key string) ([]byte, error) {
	r := GetRedis().Get()
	defer r.Close()
	res, err := redis.Bytes(r.Do("get", key))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Exists(key string) ([]byte, error) {
	r := GetRedis().Get()
	defer r.Close()
	res, err := redis.Bytes(r.Do("exists", key))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Set(key, value string) error {
	r := GetRedis().Get()
	defer r.Close()
	_, err := redis.Bytes(r.Do("set", key, value))
	if err != nil {
		return err
	}
	return nil

}

func Expire(key string,ttl int) error  {
	r := GetRedis().Get()
	defer r.Close()
	_,err := redis.Bytes(r.Do("expire",key,ttl))
	if err != nil {
		return err
	}
	return nil
}
