package handlers

import (
	"context"
	"database/sql"
	"errors"
	"github.com/danielgtaylor/huma/v2"
	i "tech-support/internal/api/handlers/input"
	o "tech-support/internal/api/handlers/output"
)

// GetHandler Generic to handle different types of Get operations
func GetHandler[DBEntity any, Output any](ctx context.Context,
	input *i.IDInput,
	dbGet func(ctx context.Context, id int64) (DBEntity, error),
	conv func(source DBEntity) Output,
) (*Output, error) {
	// Fetch the entity using the provided dbGet function
	entity, err := dbGet(ctx, input.ID)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, huma.Error400BadRequest("Object does not exist")
		}
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	// Convert the entity to the desired output format using the provided conversion function
	resp := conv(entity)
	return &resp, nil
}

// ListHandler Generic to handle different types of Get operations
func ListHandler[DBEntity any, Output any](ctx context.Context,
	dbList func(ctx context.Context) ([]DBEntity, error),
	conv func(source []DBEntity) Output,
) (*Output, error) {
	// Fetch the entity using the provided dbGet function
	entity, err := dbList(ctx)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, huma.Error400BadRequest("Object does not exist")
		}
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	// Convert the entity to the desired output format using the provided conversion function
	resp := conv(entity)
	return &resp, nil
}

// PostHandler Generic to handle different types of Post operations
func PostHandler[Params, DBEntity, Output any](ctx context.Context,
	params Params,
	dbCreate func(ctx context.Context, params Params) (DBEntity, error),
	conv func(source DBEntity) Output,
) (*Output, error) {
	// Execute database create object function
	entity, err := dbCreate(ctx, params)
	if err != nil {
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	// Convert the entity to the desired output format using the provided conversion function
	resp := conv(entity)
	return &resp, nil
}

// DeleteHandler Generic to handle different types of Delete operations
func DeleteHandler(ctx context.Context,
	input *i.IDInput,
	dbDelete func(ctx context.Context, id int64) error,
) (*o.SuccessOutput, error) {
	resp := o.SuccessOutput{}

	// Execute object deletion
	err := dbDelete(ctx, input.ID)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, huma.Error400BadRequest("Ticket does not exist")
		}
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp.Body.Success = true
	return &resp, nil
}
