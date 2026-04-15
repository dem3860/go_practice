package schema

import "go_practice/domain/entity"

type CreateUserReq struct {
	Body CreateUserReqBody
}

type CreateUserReqBody struct {
	Name     string `json:"name" maxLength:"50" example:"John Doe" doc:"user's full name"`
	Role     string `json:"role" example:"user" doc:"user's role"`
	Email    string `json:"email" format:"email" example:"john.doe@example.com" doc:"user's email address"`
	Password string `json:"password" minLength:"8" example:"password123" doc:"user's password"`
}

type CreateUserRes struct {
	Body CreateUserResBody
}

type CreateUserResBody struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func ToCreateUserResponse(user entity.User) *CreateUserRes {
	return &CreateUserRes{
		Body: CreateUserResBody{
			ID:    user.ID,
			Name:  user.Name,
			Role:  string(user.Role),
			Email: user.Email,
		},
	}
}
