package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	srvUsers "github.com/rtemb/srv-users/pkg/client/srv-users"
	"github.com/sirupsen/logrus"
)

func (h *Handler) AddRole(ctx context.Context, req *srvUsers.AddRoleRequest) (*empty.Empty, error) {
	h.logger.WithFields(logrus.Fields{"method": "handler.AddRole"}).Trace(req)
	err := h.service.AddRole(ctx, req.Uuid, req.Role)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
