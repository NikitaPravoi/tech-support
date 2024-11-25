// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Project struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Session struct {
	ID        int64              `json:"id"`
	UserID    int64              `json:"user_id"`
	Token     string             `json:"token"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
}

type SupportTicket struct {
	ID          int64              `json:"id"`
	ProjectID   int64              `json:"project_id"`
	StatusID    int64              `json:"status_id"`
	Title       string             `json:"title"`
	Description *string            `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	CreatedBy   string             `json:"created_by"`
	Answer      *string            `json:"answer"`
	AnsweredBy  *int64             `json:"answered_by"`
}

type TicketStatus struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type User struct {
	ID       int64  `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int64  `json:"role_id"`
}

type UserProject struct {
	ID        int64 `json:"id"`
	UserID    int64 `json:"user_id"`
	ProjectID int64 `json:"project_id"`
}