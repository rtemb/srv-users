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

const (
	EmailToUserIDSet = "emailToUserID:"
	UsersSet         = "users:"
)

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

func (r *RedisStorage) StoreUser(user *User) error {
	if user.ID == "" {
		return errors.Wrap(UnableToStoreUser, "user.ID is not set")
	}

	conn := r.pool.Get()
	err := conn.Send("MULTI")
	if err != nil {
		return errors.Wrap(err, StorageError.Error())
	}
	err = conn.Send("SET", EmailToUserIDSet+user.Email, user.ID)
	if err != nil {
		return errors.Wrap(err, StorageError.Error())
	}

	data, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "can't store entry in storage")
	}
	err = conn.Send("SET", UsersSet+user.ID, data)
	if err != nil {
		return errors.Wrap(err, StorageError.Error())
	}

	_, err = conn.Do("EXEC")
	if err != nil {
		return errors.Wrap(err, StorageError.Error())
	}

	return nil
}

func (r *RedisStorage) GetUserByEmail(email string) (*User, error) {
	conn := r.pool.Get()
	k, err := conn.Do("GET", EmailToUserIDSet+email)
	if k == nil {
		return nil, nil
	}

	uID := string(k.([]byte))
	data, err := conn.Do("GET", UsersSet+uID)
	if err != nil {
		return nil, errors.Wrap(err, "can't get user from storage")
	}
	if data == nil {
		return nil, nil
	}

	user := &User{}
	err = json.Unmarshal(data.([]byte), &user)
	if err != nil {
		return nil, errors.Wrap(err, "can't unmarshal user from storage")
	}

	return user, nil
}

func (r *RedisStorage) GetUserByUUID(uuid string) (*User, error) {
	conn := r.pool.Get()
	data, err := conn.Do("GET", UsersSet+uuid)
	if err != nil {
		return nil, errors.Wrap(err, "can't get user from storage")
	}
	if data == nil {
		return nil, nil
	}

	user := &User{}
	err = json.Unmarshal(data.([]byte), &user)
	if err != nil {
		return nil, errors.Wrap(err, "can't unmarshal user from storage")
	}

	return user, nil
}

func (r *RedisStorage) Set(key string, val interface{}) error {
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
