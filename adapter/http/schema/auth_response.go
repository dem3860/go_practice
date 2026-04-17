package schema

import (
	"go_practice/domain/entity"
	"time"
)

const TokenType = "Bearer"

type LoginResBody struct {
	AccessToken string      `json:"accessToken"`
	TokenType   string      `json:"tokenType"`
	User        UserResBody `json:"user"`
}

type LoginRes struct {
	Body LoginResBody
}

func ToLoginResponse(user entity.User, accessToken string) *LoginRes {
	return &LoginRes{
		Body: LoginResBody{
			AccessToken: accessToken,
			TokenType:   "Bearer",
			User:        toUserResBody(user),
		},
	}
}

type SignupResBody struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SignupRes struct {
	Body SignupResBody
}

func ToSignupResponse(user entity.User) *SignupRes {
	return &SignupRes{
		Body: SignupResBody{
			ID:        user.ID,
			Name:      user.Name,
			Role:      string(user.Role),
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
}
