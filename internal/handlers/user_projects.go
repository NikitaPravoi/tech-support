package handlers

import (
	"context"
	"database/sql"
	"errors"
	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"tech-support/internal/api/converter"
	"tech-support/internal/api/converter/generated"
	i "tech-support/internal/api/handlers/input"
	o "tech-support/internal/api/handlers/output"
	"tech-support/internal/database"
)

func ListUserProjectHandler(ctx context.Context, input *struct{}) (*o.ProjectsOutput, error) {
	userProjects, err := database.Repository.ListUserProjects(ctx, ctx.Value("user_id").(int64))
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, huma.Error400BadRequest("Projects does not exist")
		}
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp := converter.ConvertProjectsToProjectsOutput(userProjects)
	return &resp, nil
}

func CreateUserProjectHandler(ctx context.Context, input *i.CreateUserProjectInput) (*o.UserProjectOutput, error) {
	userProject, err := database.Repository.CreateUserProject(ctx, database.CreateUserProjectParams(input.Body))
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			return nil, huma.Error400BadRequest("User or project doesn't exist", err)
		}

		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, huma.Error400BadRequest("User with the same binding already exists", err)
		}
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp := generated.ConvertUserProjectToUserProjectOutput(userProject)
	return &resp, nil
}

func DeleteUserProjectHandler(ctx context.Context, input *struct {
	ID int64 `json:"id" example:"1" doc:"Project ID" path:"id"`
}) (*o.SuccessOutput, error) {
	resp := o.SuccessOutput{}

	params := database.DeleteUserProjectParams{UserID: ctx.Value("user_id").(int64), ProjectID: input.ID}

	err := database.Repository.DeleteUserProject(ctx, params)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, huma.Error400BadRequest("Ticket does not exist")
		}
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp.Body.Success = true
	return &resp, nil
}
