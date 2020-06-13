package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// nolint
var (
	UnableToCreateUser = status.Error(codes.Internal, "unable to create user")
	UserAlreadyExists  = status.Error(codes.FailedPrecondition, "user already exist")
	UnableToAddRole    = status.Error(codes.Internal, "unable to add role")
	UserNotFound       = status.Error(codes.NotFound, "unable to find user")
)
