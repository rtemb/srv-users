package token_encoder

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rtemb/srv-users/internal/storage"
	"github.com/rtemb/srv-users/pkg/token_decoder"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -fake-name TokenEncoderMock -o ../testing/mocks/token_encoder.go . AuthableEncoder
type AuthableEncoder interface {
	Encode(user *storage.User) (string, error)
}

type TokenEncoder struct {
	key []byte
}

func NewEncoder(key string) *TokenEncoder {
	return &TokenEncoder{key: []byte(key)}
}

// Encode a claim into a JWT
func (e *TokenEncoder) Encode(user *storage.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Store the Claims
	claims := token_decoder.CustomClaims{
		User: &token_decoder.JWTUser{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "shippy.user",
		},
	}

	// Store token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(e.key)
}
