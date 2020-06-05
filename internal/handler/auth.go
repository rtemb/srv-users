package handler

import (
	"context"

	"github.com/pkg/errors"
	srvUsers "github.com/rtemb/srv-users/pkg/client/srv-users"
)

func (h *Handler) Auth(ctx context.Context, req *srvUsers.AuthRequest) (*srvUsers.AuthResponse, error) {
	r := &srvUsers.AuthResponse{}

	token, err := h.service.Auth(ctx, req.Email, req.Password)
	if err != nil {
		return r, errors.Wrap(err, "unable to authenticate")
	}

	r.Token = token
	r.Valid = true
	return r, nil
}
