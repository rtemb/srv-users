package storage_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/rtemb/srv-users/internal/config"
	store "github.com/rtemb/srv-users/internal/storage"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type RedisSuite struct {
	suite.Suite
	store  store.Storage
	Logger *logrus.Entry
}

func (a *RedisSuite) SetupSuite() {
	err := os.Setenv("TOKEN_KEY", "test")
	a.Require().NoError(err)

	cfg, err := config.Load()
	a.Require().NoError(err)

	a.store = store.NewStorage(cfg.Redis)
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, &RedisSuite{})
}

func (a *RedisSuite) Test_StoreAndGet() {
	user := &store.User{
		Name:     "test-name",
		Company:  "test-company",
		Email:    "test@example.com",
		Password: "test-pass",
	}
	err := a.store.Store(user.Email, user)
	a.NoError(err)

	res, err := a.store.Get(user.Email)
	a.Require().NoError(err)
	a.NotNil(res)

	u := &store.User{}
	err = json.Unmarshal(res.([]byte), &u)
	a.Require().NoError(err)

	a.Equal(user.Name, u.Name)
	a.Equal(user.Company, u.Company)
	a.Equal(user.Email, u.Email)
	a.Equal(user.Password, u.Password)
}
