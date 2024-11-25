package input

type CreateUserInput struct {
	Body struct {
		Login    string `json:"login" maxLength:"50" doc:"User login"`
		Email    string `json:"email" format:"email" example:"admin@admin.com" doc:"User email"`
		Password string `json:"password" maxLength:"50" doc:"User password"`
		RoleID   int64  `json:"role_id" example:"1" doc:"User role id"`
	}
}

type LoginInput struct {
	Body struct {
		Login    string `json:"login" example:"admin" doc:"User login"`
		Password string `json:"password" example:"admin" doc:"User password"`
	}
}
