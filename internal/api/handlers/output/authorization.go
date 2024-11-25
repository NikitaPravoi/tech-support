package output

import (
	"time"
)

type User struct {
	ID       int64  `json:"id" example:"1" doc:"User id"`
	Login    string `json:"login" example:"admin" doc:"User login"`
	Email    string `json:"email" format:"email" example:"admin@admin.com" doc:"User email"`
	Password string `json:"password" example:"User password" doc:"User password"`
	RoleID   int64  `json:"role_id" example:"1" doc:"User role id"`
}

type CreateUserOutput struct {
	Body User
}

type Session struct {
	ID        int64     `json:"id" example:"1" doc:"Session ID"`
	UserID    int64     `json:"user_id" example:"1" doc:"User ID"`
	Token     string    `json:"token" example:"1430686b877ab3d9e5ad4dff66b16e62f378dabd96d3ab7b6f00e2b5e1ff47c8" doc:"Token"`
	CreatedAt time.Time `json:"created_at" format:"date-time" doc:"Create Date"`
	ExpiresAt time.Time `json:"expires_at" format:"date-time" doc:"Expires Date"`
}

type LoginOutput struct {
	Body Session
}
