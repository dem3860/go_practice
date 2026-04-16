package output

import "go_practice/domain/entity"

type TokenProvider interface {
	GenerateToken(user entity.User) (string, error)
	ValidateToken(token string) (string, error)
}
