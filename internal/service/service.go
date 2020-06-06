package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rtemb/srv-users/internal/storage"
	"github.com/rtemb/srv-users/internal/token_encoder"
	srvErr "github.com/rtemb/srv-users/pkg/client/errors"
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
	if err = s.store.AddUser(&user); err != nil {
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
