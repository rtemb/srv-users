package storage

import (
	"encoding/json"
	"net"

	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"github.com/rtemb/srv-users/internal/config"
)

type RedisStorage struct {
	pool *redis.Pool
}

var _ Storage = (*RedisStorage)(nil)

func NewStorage(cfg *config.Redis) *RedisStorage {
	redisPool := &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		IdleTimeout: cfg.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			opts := []redis.DialOption{
				redis.DialDatabase(0),
			}

			if cfg.Password != "" {
				opts = append(opts, redis.DialPassword(cfg.Password))
			}
			return redis.Dial("tcp", net.JoinHostPort(cfg.Host, cfg.Port), opts...)
		},
	}

	return &RedisStorage{pool: redisPool}
}

func (r *RedisStorage) Store(key string, val interface{}) error {
	j, err := json.Marshal(val)
	if err != nil {
		return errors.Wrap(err, "can't store entry in storage")
	}

	_, err = r.pool.Get().Do("SET", key, j)
	if err != nil {
		return errors.Wrap(err, "can't store entry in storage")
	}

	return nil
}

func (r *RedisStorage) Get(key string) (interface{}, error) {
	rsp, err := r.pool.Get().Do("GET", key)
	if err != nil {
		return nil, errors.Wrap(err, "can't get entry from storage")
	}
	return rsp, nil
}

func (r *RedisStorage) GetUserByEmail(email string) (*User, error) {
	data, err := r.Get(email)
	if err != nil {
		return nil, errors.Wrap(err, "can't get user from storage")
	}
	user := &User{}
	err = json.Unmarshal(data.([]byte), &user)
	if err != nil {
		return nil, errors.Wrap(err, "can't unmarshal user from storage")
	}

	return user, nil
}
