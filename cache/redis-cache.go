package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// var ctx = context.Background()

type RedisCache struct {
	Host     string
	Password string
	Db       int
	Duration time.Duration
	Context  context.Context
}

func NewRedisCache(host string, password string, duration time.Duration) *RedisCache {
	return &RedisCache{
		Host:     host,
		Password: password,
		Duration: duration,
		Context:  context.Background(),
	}
}

func (cache *RedisCache) createClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.Host,
		Password: cache.Password,
		DB:       cache.Db,
	})
}

func (cache *RedisCache) SetValue(key string, val interface{}) {
	err := cache.createClient().
		Set(
			cache.Context,
			key,
			fmt.Sprintf("%v", val),
			cache.Duration,
		).
		Err()

	if err != nil {
		panic(err)
	}
}

func (cache *RedisCache) AppendValue(key string, val interface{}) {
	err := cache.createClient().
		Append(
			cache.Context,
			key,
			fmt.Sprintf("%v", val),
		).
		Err()

	if err != nil {
		panic(err)
	}
}

func (cache *RedisCache) GetValue(key string) interface{} {
	val, err := cache.createClient().
		Get(
			cache.Context,
			key,
		).
		Result()

	if err == redis.Nil {
		return nil
	}
	if err != nil {
		panic(err)
	}

	return val
}
