package storage

import "errors"

var (
	UnableToCreateUser = errors.New("could not create a user")
	DataStoreError     = errors.New("can't store entry in storage")
)
