package storage

import srvUsers "github.com/rtemb/srv-users/pkg/client/srv-users"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -fake-name StorageMock -o ../testing/mocks/storage.go . Storage
type Storage interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByUUID(uuid string) (*User, error)
	StoreUser(user *User) error
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -fake-name SimpleStorageMock -o ../testing/mocks/simple_storage.go . SimpleStorage
type SimpleStorage interface {
	Set(key string, val interface{}) error
	Get(key string) (interface{}, error)
}

type User struct {
	ID       string
	Name     string
	Company  string
	Email    string
	Password string
	Roles    map[srvUsers.Role]struct{}
}
