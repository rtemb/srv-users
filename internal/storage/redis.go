package storage

import (
	"encoding/json"
	"net"

	"github.com/garyburd/redigo/redis"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rtemb/srv-users/internal/config"
)

type RedisStorage struct {
	pool *redis.Pool
}

var _ Storage = (*RedisStorage)(nil)

var emailToUserIDSetName = "emailToUserIDSet:"

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

func (r *RedisStorage) AddUser(user *User) error {
	existUser, err := r.Get(user.Email)
	if err != nil {
		return errors.Wrap(err, UnableToCreateUser.Error())
	}
	if existUser != nil {
		return errors.New(UnableToCreateUser.Error())
	}

	user.ID = uuid.New().String()

	conn := r.pool.Get()
	err = conn.Send("MULTI")
	if err != nil {
		return errors.Wrap(err, DataStoreError.Error())
	}
	err = conn.Send("SET", emailToUserIDSetName+user.Email, user.ID)
	if err != nil {
		return errors.Wrap(err, DataStoreError.Error())
	}

	data, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "can't store entry in storage")
	}
	err = conn.Send("SET", user.ID, data)
	if err != nil {
		return errors.Wrap(err, DataStoreError.Error())
	}

	_, err = conn.Do("EXEC")
	if err != nil {
		return errors.Wrap(err, DataStoreError.Error())
	}

	return nil
}

func (r *RedisStorage) GetUserByEmail(email string) (*User, error) {
	conn := r.pool.Get()
	uID, err := conn.Do("GET", emailToUserIDSetName+email)
	data, err := conn.Do("GET", uID)
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
