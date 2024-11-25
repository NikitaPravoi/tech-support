package handlers

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"tech-support/internal/api/converter/generated"
	i "tech-support/internal/api/handlers/input"
	o "tech-support/internal/api/handlers/output"
	"tech-support/internal/database"
	"tech-support/internal/services"
	"time"
)

func CreateUserHandler(ctx context.Context, input *i.CreateUserInput) (*o.CreateUserOutput, error) {
	var err error

	input.Body.Password, err = services.HashPassword(input.Body.Password)
	if err != nil {
		return nil, huma.Error500InternalServerError("Hashing password error", err)
	}

	params := database.CreateUserParams(input.Body)

	user, err := database.Repository.CreateUser(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, huma.Error400BadRequest("User with the same login or email already exists", err)
		}

		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp := o.CreateUserOutput{}
	resp.Body = struct {
		ID       int64  `json:"id" example:"1" doc:"User id"`
		Login    string `json:"login" example:"admin" doc:"User login"`
		Email    string `json:"email" format:"email" example:"admin@admin.com" doc:"User email"`
		Password string `json:"password" example:"User password" doc:"User password"`
		RoleID   int64  `json:"role_id" example:"1" doc:"User role id"`
	}(user)

	return &resp, nil
}

func LoginHandler(ctx context.Context, input *i.LoginInput) (*o.LoginOutput, error) {
	resp := o.LoginOutput{}

	userCredentials, err := database.Repository.GetUserCredentials(ctx, input.Body.Login)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return &resp, huma.Error400BadRequest("User is not registered")
		}
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	if res := services.CheckPassword(userCredentials.Password, input.Body.Password); !res {
		return &resp, huma.Error403Forbidden("Incorrect password")
	}

	token := sha256.New()
	token.Write([]byte(time.Now().String()))
	hash := hex.EncodeToString(token.Sum(nil))

	params := database.CreateSessionParams{
		UserID: userCredentials.ID,
		Token:  hash,
	}

	session, err := database.Repository.CreateSession(ctx, params)
	if err != nil {
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp = generated.ConvertSessionToLoginOutput(session)

	return &resp, nil
}

func LogoutHandler(ctx context.Context, input *struct{}) (*o.SuccessOutput, error) {
	resp := &o.SuccessOutput{}

	err := database.Repository.DeleteSessionsByUserID(ctx, ctx.Value("user_id").(int64))
	if err != nil {
		resp.Body.Success = false
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp.Body.Success = true
	return resp, nil
}
