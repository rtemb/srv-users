package token_decoder

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type JWTUser struct {
	ID    string
	Name  string
	Email string
}

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *JWTUser
	jwt.StandardClaims
}

type AuthableDecoder interface {
	Decode(token string) (*CustomClaims, error)
}

type TokenDecoder struct {
	key []byte
}

func NewDecoder(key string) *TokenDecoder {
	return &TokenDecoder{key: []byte(key)}
}

// Decode a token string into a token object
func (d *TokenDecoder) Decode(tokenString string) (*CustomClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return d.key, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse token")
	}

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
