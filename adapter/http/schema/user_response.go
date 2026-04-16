package schema

import "go_practice/domain/entity"

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
		items = append(items, UserResBody{
			ID:    user.ID,
			Name:  user.Name,
			Role:  string(user.Role),
			Email: user.Email,
		})
	}

	return &ListUsersRes{
		Body: ListUsersResBody{
			Data:     items,
			Total:    total,
			NextPage: nextPage,
		},
	}
}
