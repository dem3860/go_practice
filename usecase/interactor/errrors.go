package interactor

import "errors"

var ErrKind struct {
	NotFound   error
	Conflict   error
	BadRequest error
	Validation error
	DB    error
	Hash error
	InternalServerError error
}

func init() {
	ErrKind.NotFound = errors.New("not found error")
	ErrKind.Conflict = errors.New("conflict error")
	ErrKind.BadRequest = errors.New("bad request error")
	ErrKind.Validation = errors.New("validation error")
	ErrKind.DB = errors.New("database error")
	ErrKind.Hash = errors.New("hashing error")
	ErrKind.InternalServerError = errors.New("internal server error")
}
