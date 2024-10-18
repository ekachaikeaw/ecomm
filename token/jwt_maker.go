package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) *JWTMaker {
	return &JWTMaker{secretKey: secretKey}
}

func (maker *JWTMaker) CreateToken(id int64, email string, isAdmin bool, duration time.Duration) (string, *UserClaims, error) {
	jwt.NewWithClaims(jwt.SigningMethodES256, )
}