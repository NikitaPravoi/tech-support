package handlers

import (
	"context"
	"github.com/danielgtaylor/huma/v2"
	"tech-support/internal/api/converter"
	"tech-support/internal/api/converter/generated"
	i "tech-support/internal/api/handlers/input"
	o "tech-support/internal/api/handlers/output"
	"tech-support/internal/database"
)

func GetProjectHandler(ctx context.Context, input *i.IDInput) (*o.ProjectOutput, error) {
	return GetHandler(
		ctx,
		input,
		database.Repository.GetProject,
		generated.ConvertProjectToProjectOutput)
}

func ListProjectHandler(ctx context.Context, input *struct{}) (*o.ProjectsOutput, error) {
	return ListHandler(
		ctx,
		database.Repository.ListProjects,
		converter.ConvertProjectsToProjectsOutput)
}

func CreateProjectHandler(ctx context.Context, input *i.CreateProjectInput) (*o.ProjectOutput, error) {
	return PostHandler(
		ctx,
		database.CreateProjectParams(input.Body),
		database.Repository.CreateProject,
		generated.ConvertProjectToProjectOutput)
}

func UpdateProjectHandler(ctx context.Context, input *i.UpdateProjectInput) (*o.ProjectOutput, error) {
	params := database.UpdateProjectParams{
		ID:          input.ID,
		Name:        input.Body.Name,
		Description: input.Body.Description,
	}

	project, err := database.Repository.UpdateProject(ctx, params)
	if err != nil {
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp := generated.ConvertProjectToProjectOutput(project)
	return &resp, nil
}

func DeleteProjectHandler(ctx context.Context, input *i.IDInput) (*o.SuccessOutput, error) {
	return DeleteHandler(ctx, input, database.Repository.DeleteProject)
}
