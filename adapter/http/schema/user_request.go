package schema

import inputport "go_practice/usecase/port/input"

type ListUsersReq struct {
	Page     int    `query:"page" minimum:"1" example:"1" doc:"page number"`
	Take     int    `query:"take" minimum:"1" maximum:"100" example:"20" doc:"number of users to return"`
	Q        string `query:"q" example:"yamada" doc:"search keyword for name or email"`
	Order    string `query:"order" enum:"asc,desc" example:"desc" doc:"sort order"`
	OrderBy  string `query:"orderBy" enum:"createdAt,name,email,role" example:"createdAt" doc:"field to sort by"`
	UserType string `query:"userType" enum:"user,admin" example:"user" doc:"user role filter"`
}

func (r *ListUsersReq) ToQuery() inputport.ListUsersQuery {
	return inputport.ListUsersQuery{
		Page:     firstNonZero(r.Page, 1),
		Take:     firstNonZero(r.Take, 20),
		Q:        r.Q,
		Order:    firstNonEmpty(r.Order, "desc"),
		OrderBy:  firstNonEmpty(r.OrderBy, "createdAt"),
		UserType: r.UserType,
	}
}

func firstNonZero(value, fallback int) int {
	if value == 0 {
		return fallback
	}

	return value
}

func firstNonEmpty(value, fallback string) string {
	if value == "" {
		return fallback
	}

	return value
}

type UpdateByMeReqBody struct {
	Name string `json:"name" maxLength:"50" example:"John Doe" doc:"user's full name"`
}

type UpdateByMeReq struct {
	Body UpdateByMeReqBody
}

type DeleteUserReq struct {
	UserID string `path:"userID" doc:"user ID to delete" example:"01HXYZ1234567890ABCDEFGHJK"`
}
