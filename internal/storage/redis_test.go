package storage_test

import (
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/rtemb/srv-users/internal/config"
	store "github.com/rtemb/srv-users/internal/storage"
	srvUsers "github.com/rtemb/srv-users/pkg/client/srv-users"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type RedisSuite struct {
	suite.Suite
	store  *store.RedisStorage
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
	err := a.store.Set(user.Email, user)
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

func (a *RedisSuite) Test_AddUser() {
	user := &store.User{
		ID:       uuid.New().String(),
		Name:     "test-name",
		Company:  "test-company",
		Email:    "test" + strconv.Itoa(rand.Intn(10000)) + "@example.com",
		Password: "test-pass",
	}
	err := a.store.StoreUser(user)
	a.NoError(err)

	u, err := a.store.GetUserByEmail(user.Email)
	a.Require().NoError(err)
	a.NotNil(u)

	a.Equal(user.Name, u.Name)
	a.Equal(user.Company, u.Company)
	a.Equal(user.Email, u.Email)
	a.Equal(user.Password, u.Password)
}

func (a *RedisSuite) Test_AddRole() {
	role := srvUsers.Role_USER
	user := &store.User{
		ID:       uuid.New().String(),
		Name:     "test-name",
		Company:  "test-company",
		Email:    "test" + strconv.Itoa(rand.Intn(10000)) + "@example.com",
		Password: "test-pass",
		Roles: map[srvUsers.Role]struct{}{
			role: {},
		},
	}
	err := a.store.StoreUser(user)
	a.NoError(err)

	u, err := a.store.GetUserByUUID(user.ID)
	a.Require().NoError(err)
	a.Require().NotNil(u)

	u.Roles[srvUsers.Role_USER_ADMIN] = struct{}{}
	err = a.store.StoreUser(user)
	a.Require().NoError(err)

	updatedUser, err := a.store.GetUserByUUID(user.ID)
	a.Require().NoError(err)
	a.NotNil(u)

	a.Equal(user.Name, updatedUser.Name)
	a.Equal(user.Company, updatedUser.Company)
	a.Equal(user.Email, updatedUser.Email)
	a.Equal(user.Password, updatedUser.Password)
	a.NotNil(user.Roles[srvUsers.Role_USER_ADMIN])
	a.NotNil(user.Roles[srvUsers.Role_USER])
}
