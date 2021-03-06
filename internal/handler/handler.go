package handler

import (
	"context"

	"github.com/rtemb/srv-users/internal/storage"
	srvUsers "github.com/rtemb/srv-users/pkg/client/srv-users"
	"github.com/sirupsen/logrus"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 --fake-name ServiceHandlerMock -o ../testing/mocks/service_handler.go . ServiceHandler
type ServiceHandler interface {
	CreateUser(ctx context.Context, user storage.User) error
	Auth(ctx context.Context, email string, pass string) (string, error)
	AddRole(ctx context.Context, uuid string, role srvUsers.Role) error
}

type Handler struct {
	service ServiceHandler
	logger  *logrus.Entry
}

func NewHandler(s ServiceHandler, l *logrus.Entry) *Handler {
	return &Handler{service: s, logger: l}
}
