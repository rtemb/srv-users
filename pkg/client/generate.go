//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -fake-name srvUsersMock -o srv-users/mocks/mocks.go ./srv-users UsersServiceClient

package client
