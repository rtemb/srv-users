module github.com/rtemb/srv-users

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/garyburd/redigo v1.6.0
	github.com/google/uuid v1.1.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/rtemb/srv-users/pkg/client v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.0
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	google.golang.org/grpc v1.29.1
)

replace github.com/rtemb/srv-users/pkg/client => ./pkg/client
