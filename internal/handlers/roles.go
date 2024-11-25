package handlers

import (
	"context"
	"database/sql"
	"errors"
	"github.com/danielgtaylor/huma/v2"
	"tech-support/internal/api/converter"
	"tech-support/internal/api/converter/generated"
	i "tech-support/internal/api/handlers/input"
	o "tech-support/internal/api/handlers/output"
	"tech-support/internal/database"
)

func GetRoleHandler(ctx context.Context, input *i.IDInput) (*o.RoleOutput, error) {
	return GetHandler(
		ctx,
		input,
		database.Repository.GetRole,
		generated.ConvertRoleToRoleOutput)
}

func GetUserRoleHandler(ctx context.Context, input *struct{}) (*o.GetUserRoleOutput, error) {
	userRole, err := database.Repository.GetUserWithHisRoles(ctx, ctx.Value("user_id").(int64))
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, huma.Error400BadRequest("User does not exist")
		}
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp := o.GetUserRoleOutput{Body: o.UserWithRoleInfo(userRole)}

	return &resp, nil
}

func ListRoleHandler(ctx context.Context, input *struct{}) (*o.RolesOutput, error) {
	return ListHandler(
		ctx,
		database.Repository.ListRoles,
		converter.ConvertRolesToRolesOutput)
}

func CreateRoleHandler(ctx context.Context, input *i.CreateRoleInput) (*o.RoleOutput, error) {
	return PostHandler(
		ctx,
		database.CreateRoleParams(input.Body),
		database.Repository.CreateRole,
		generated.ConvertRoleToRoleOutput)
}

func UpdateRoleHandler(ctx context.Context, input *i.UpdateRoleInput) (*o.RoleOutput, error) {
	params := database.UpdateRoleParams{
		ID:          input.ID,
		Name:        input.Body.Name,
		Description: input.Body.Description,
	}

	role, err := database.Repository.UpdateRole(ctx, params)
	if err != nil {
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp := generated.ConvertRoleToRoleOutput(role)
	return &resp, nil
}

func DeleteRoleHandler(ctx context.Context, input *i.IDInput) (*o.SuccessOutput, error) {
	return DeleteHandler(ctx, input, database.Repository.DeleteRole)
}
