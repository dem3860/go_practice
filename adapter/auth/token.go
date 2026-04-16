package auth

import (
	"fmt"
	"time"

	"go_practice/domain/entity"

	"github.com/golang-jwt/jwt/v5"
)

type authClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

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
	now := time.Now()
	claims := authClaims{
		Role: string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID,
			ExpiresAt: jwt.NewNumericDate(now.Add(a.expire)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.secretKey)
}

func (a *JWTProvider) ValidateToken(token string) (string, error) {
	claims := &authClaims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return a.secretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return "", err
	}
	if !parsed.Valid {
		return "", fmt.Errorf("token is invalid")
	}
	if claims.Subject == "" {
		return "", fmt.Errorf("token subject is missing")
	}

	return claims.Subject, nil
}
