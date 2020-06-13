package storage

import "errors"

var (
	UnableToStoreUser = errors.New("could not create a user")
	StorageError      = errors.New("can't store entry in storage")
)
