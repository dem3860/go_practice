package schema

import (
	"go_practice/domain/entity"
	"time"
)

type UserResBody struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserRes struct {
	Body UserResBody
}

func toUserResBody(user entity.User) UserResBody {
	return UserResBody{
		ID:        user.ID,
		Name:      user.Name,
		Role:      string(user.Role),
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponse(user entity.User) *UserRes {
	return &UserRes{
		Body: toUserResBody(user),
	}
}

// nextPageはポインタ(次ページ無しの場合はnilの可能性がある)
type ListUsersResBody struct {
	Data     []UserResBody `json:"data"`
	Total    int           `json:"total"`
	NextPage *int          `json:"nextPage"`
}

type ListUsersRes struct {
	Body ListUsersResBody
}

func ToListUsersResponse(users []entity.User, total int, nextPage *int) *ListUsersRes {
	items := make([]UserResBody, 0, len(users))
	for _, user := range users {
		items = append(items, toUserResBody(user))
	}

	return &ListUsersRes{
		Body: ListUsersResBody{
			Data:     items,
			Total:    total,
			NextPage: nextPage,
		},
	}
}
