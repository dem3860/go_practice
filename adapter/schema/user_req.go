package schema

type LoginReqBody struct {
	Email    string `json:"email" format:"email" example:"john.doe@example.com"`
	Password string `json:"password" minLength:"8" example:"password123"`
}

type LoginReq struct {
	Body LoginReqBody
}

type CreateUserReqBody struct {
	Name     string `json:"name" maxLength:"50" example:"John Doe" doc:"user's full name"`
	Role     string `json:"role" example:"user" doc:"user's role"`
	Email    string `json:"email" format:"email" example:"john.doe@example.com" doc:"user's email address"`
	Password string `json:"password" minLength:"8" example:"password123" doc:"user's password"`
}

type CreateUserReq struct {
	Body CreateUserReqBody
}
