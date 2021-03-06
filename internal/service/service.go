package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rtemb/srv-users/internal/storage"
	"github.com/rtemb/srv-users/internal/token_encoder"
	srvErr "github.com/rtemb/srv-users/pkg/client/errors"
	srvUsers "github.com/rtemb/srv-users/pkg/client/srv-users"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	logger       *logrus.Entry
	store        storage.Storage
	tokenEncoder token_encoder.AuthableEncoder
}

func NewService(l *logrus.Entry, s storage.Storage, te token_encoder.AuthableEncoder) *Service {
	return &Service{logger: l, store: s, tokenEncoder: te}
}

func (s *Service) CreateUser(ctx context.Context, user storage.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "error hashing password")
	}

	exist, err := s.store.GetUserByEmail(user.Email)
	if err != nil {
		return errors.Wrap(err, srvErr.UnableToCreateUser.Error())
	}
	if exist != nil {
		return srvErr.UserAlreadyExists
	}

	user.Password = string(hashedPass)
	user.ID = uuid.New().String()
	user.Roles = make(map[srvUsers.Role]struct{})

	if err = s.store.StoreUser(&user); err != nil {
		return errors.Wrap(err, srvErr.UnableToCreateUser.Error())
	}

	return nil
}

func (s *Service) Auth(ctx context.Context, email string, pass string) (string, error) {
	token := ""
	user, err := s.store.GetUserByEmail(email)
	if err != nil {
		return token, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return token, err
	}

	token, err = s.tokenEncoder.Encode(user)
	if err != nil {
		return token, err
	}

	return token, err
}

func (s *Service) AddRole(ctx context.Context, uuid string, role srvUsers.Role) error {
	user, err := s.store.GetUserByUUID(uuid)
	if err != nil {
		return errors.Wrap(srvErr.UnableToAddRole, err.Error())
	}
	if user == nil {
		return srvErr.UserNotFound
	}

	user.Roles[role] = struct{}{}

	err = s.store.StoreUser(user)
	if err != nil {
		return errors.Wrap(srvErr.UnableToAddRole, err.Error())
	}

	return nil
}
