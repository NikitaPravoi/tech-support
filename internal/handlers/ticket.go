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

func GetTicketHandler(ctx context.Context, input *i.IDInput) (*o.SupportTicketOutput, error) {
	return GetHandler(
		ctx,
		input,
		database.Repository.GetSupportTicket,
		generated.ConvertSupportTicketToSupportTicketOutput)
}

func ListTicketHandler(ctx context.Context, input *struct{}) (*o.SupportTicketsOutput, error) {
	return ListHandler(
		ctx,
		database.Repository.ListSupportTickets,
		converter.ConvertSupportTicketsToSupportTicketsOutput)
}

func CreateTicketHandler(ctx context.Context, input *i.CreateSupportTicketInput) (*o.SupportTicketOutput, error) {
	return PostHandler(
		ctx,
		database.CreateSupportTicketParams(input.Body),
		database.Repository.CreateSupportTicket,
		generated.ConvertSupportTicketToSupportTicketOutput)
}

func UpdateTicketHandler(ctx context.Context, input *i.UpdateSupportTicketInput) (*o.SupportTicketOutput, error) {
	params := database.UpdateSupportTicketParams{
		ID:          input.ID,
		ProjectID:   input.Body.ProjectID,
		StatusID:    input.Body.StatusID,
		Title:       input.Body.Title,
		Description: input.Body.Description,
		CreatedBy:   input.Body.CreatedBy,
	}

	supportTicket, err := database.Repository.UpdateSupportTicket(ctx, params)
	if err != nil {
		return nil, huma.Error500InternalServerError("Database error", err)
	}

	resp := generated.ConvertSupportTicketToSupportTicketOutput(supportTicket)
	return &resp, nil
}

func DeleteTicketHandler(ctx context.Context, input *i.IDInput) (*o.SuccessOutput, error) {
	return DeleteHandler(ctx, input, database.Repository.DeleteSupportTicket)
}
