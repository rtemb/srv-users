package handler

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rtemb/srv-users/internal/storage"
	srvUsers "github.com/rtemb/srv-users/pkg/client/srv-users"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateUser(ctx context.Context, req *srvUsers.CreateUserRequest) (*srvUsers.CreateUserResponse, error) {
	h.logger.WithFields(logrus.Fields{"method": "handler.CreateUser"}).Trace(req)

	rsp := &srvUsers.CreateUserResponse{}
	user := storage.User{
		Company:  req.Company,
		Email:    req.Email,
		Password: req.Password,
	}
	err := h.service.CreateUser(ctx, user)
	if err != nil {
		err = errors.Wrap(err, "unable to create user")
		h.logger.Error(err)
		return rsp, err
	}

	return rsp, nil
}
