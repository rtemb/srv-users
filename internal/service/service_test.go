package service

import (
	"context"
	"testing"

	store "github.com/rtemb/srv-users/internal/storage"
	"github.com/rtemb/srv-users/internal/testing/mocks"
	"github.com/rtemb/srv-users/internal/token_encoder"
	"github.com/rtemb/srv-users/pkg/token_decoder"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	service     *Service
	logger      *logrus.Entry
	hook        *test.Hook
	storageMock *mocks.StorageMock
	tokenKey    string
}

func (a *ServiceSuite) SetupSuite() {
	l, h := test.NewNullLogger()
	a.logger = logrus.NewEntry(l)
	a.hook = h

	a.storageMock = &mocks.StorageMock{}
	a.tokenKey = "test-1234"

	encoder := token_encoder.NewEncoder(a.tokenKey)
	a.service = NewService(a.logger, a.storageMock, encoder)
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, &ServiceSuite{})
}

func (a *ServiceSuite) Test_CreateUser() {
	user := store.User{
		Name:     "test-name",
		Company:  "test-company",
		Email:    "test@example.com",
		Password: "test-pass",
	}

	err := a.service.CreateUser(context.Background(), user)
	a.Require().NoError(err)
}

func (a *ServiceSuite) Test_Auth() {
	user := store.User{
		ID:       "qwerty1234",
		Name:     "test-name",
		Company:  "test-company",
		Email:    "test@example.com",
		Password: "$2a$10$UcSqWuknatK8L0FAZ8aOF.3S/MRuacTMoGHow2h7Knchug5Q0BOrC",
	}

	a.storageMock.GetUserByEmailCalls(func(email string) (*store.User, error) {
		a.Equal(user.Email, email)
		return &user, nil
	})

	token, err := a.service.Auth(context.Background(), "test@example.com", "test-pass")
	a.Require().NoError(err)
	a.NotEmpty(token)

	decoder := token_decoder.NewDecoder(a.tokenKey)

	claims, err := decoder.Decode(token)
	a.Require().NoError(err)
	a.NotEmpty(claims)
	a.NotEmpty(claims.StandardClaims)

	a.Equal(user.ID, claims.User.ID)
	a.Equal(user.Company, claims.User.Company)
	a.Equal(user.Email, claims.User.Email)
}
