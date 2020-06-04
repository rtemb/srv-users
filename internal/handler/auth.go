package handler

import (
	"context"

	srvUsers "github.com/rtemb/srv-users/pkg/client/srv-users"
)

func (h *Handler) Auth(ctx context.Context, req *srvUsers.AuthRequest) (*srvUsers.AuthResponse, error) {
	r := &srvUsers.AuthResponse{}
	return r, nil
}
