package db

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheStore interface {
	Get(key string) (string, error)
	Set(key string) error
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

func (st *Store) Set(key string) error {
	return nil
}

func (st *Store) Close() {
	st.Close()
}
