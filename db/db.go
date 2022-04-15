package db

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheStore interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	SetEx(key string, value interface{}, expiry time.Duration) error
	Close()
}

type Store struct {
	rdb redis.Client
}

func New() *Store {
	rdb := *redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	status := rdb.Ping(ctx)
	log.Println(status)
	return &Store{rdb: rdb}
}

func (st *Store) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	val, err := st.rdb.Get(ctx, key).Result()
	return val, err
}

func (st *Store) Set(key string, value interface{}) error {
	return nil
}

func (st *Store) SetEx(key string, value interface{}, expiry time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := st.rdb.SetEX(ctx, key, string(data), expiry).Result()
	log.Println(res)
	return err
}

func (st *Store) Close() {
	st.rdb.Close()
}
