package handler

import (
	"github.com/rtemb/srv-users/internal/storage"
	"github.com/sirupsen/logrus"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 --fake-name ServiceHandlerMock -o ../testing/mocks/service_handler.go . HandlerService
type ServiceHandler interface {
	CreateUser(user storage.User) error
	Auth(email string, pass string) (string, error)
}

type Handler struct {
	service ServiceHandler
	logger  *logrus.Entry
}

func NewHandler(s ServiceHandler, l *logrus.Entry) *Handler {
	return &Handler{service: s, logger: l}
}
