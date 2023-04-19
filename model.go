package main

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	initRedisDB()
}

func initRedisDB() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // use the default DB
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}
}

func setV(long string, short string, value string) (err error) {
	err1 := rdb.Set(ctx, long, value, time.Hour).Err()
	err0 := rdb.Set(ctx, short, long, time.Hour).Err()
	err = errors.Join(err0, err1)
	return
}

func getV(key string) (value string, err error) {
	strcmd := rdb.Get(ctx, key)
	value = strcmd.Val()
	err = strcmd.Err()
	return
}
