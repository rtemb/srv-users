//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -fake-name SrvUsersMock -o srv-users/mocks/mocks.go ./srv-users UsersServiceClient

package client
