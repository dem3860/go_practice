package auth

import (
	"time"

	"go_practice/domain/entity"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	secretKey []byte
	expire    time.Duration
}

func NewJWTProvider(secret string, expire time.Duration) *JWTProvider {
	return &JWTProvider{
		secretKey: []byte(secret),
		expire:    expire,
	}
}

func (a *JWTProvider) GenerateToken(user entity.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": string(user.Role),
		"exp":  time.Now().Add(a.expire).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.secretKey)
}
