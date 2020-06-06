package storage

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -fake-name StorageMock -o ../testing/mocks/storage.go . Storage
type Storage interface {
	Store(key string, val interface{}) error
	Get(key string) (interface{}, error)
	GetUserByEmail(email string) (*User, error)
	AddUser(user *User) error
}

type User struct {
	ID       string
	Name     string
	Company  string
	Email    string
	Password string
}
