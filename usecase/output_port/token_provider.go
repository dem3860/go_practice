package output_port

import "go_practice/domain/entity"

type TokenProvider interface {
	GenerateToken(user entity.User) (string, error)
}
