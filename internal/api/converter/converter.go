package converter

import (
	"github.com/jackc/pgx/v5/pgtype"
	"tech-support/internal/api/handlers/output"
	"tech-support/internal/database"
	"time"
)

// goverter:converter
// goverter:output:format function
// goverter:output:file ./generated/generated.go
// goverter:extend TimestamptzToTime
type Converter interface {
	// goverter:map . Body
	ConvertSessionToLoginOutput(source database.Session) output.LoginOutput
	// goverter:map . Body
	ConvertSupportTicketToSupportTicketOutput(source database.SupportTicket) output.SupportTicketOutput
	// goverter:map . Body
	ConvertProjectToProjectOutput(source database.Project) output.ProjectOutput
	// goverter:map . Body
	ConvertRoleToRoleOutput(source database.Role) output.RoleOutput
	// goverter:map . Body
	ConvertUserProjectToUserProjectOutput(source database.UserProject) output.UserProjectOutput
}

func TimestamptzToTime(timestamp pgtype.Timestamptz) time.Time {
	return timestamp.Time
}

func ConvertSupportTicketsToSupportTicketsOutput(source []database.SupportTicket) output.SupportTicketsOutput {
	out := output.SupportTicketsOutput{}

	for _, value := range source {
		out.Body = append(out.Body, output.SupportTicket{ID: value.ID,
			ProjectID:   value.ProjectID,
			StatusID:    value.StatusID,
			Title:       value.Title,
			Description: value.Description,
			CreatedAt:   value.CreatedAt.Time,
			UpdatedAt:   value.UpdatedAt.Time,
			CreatedBy:   value.CreatedBy,
			Answer:      value.Answer,
			AnsweredBy:  value.AnsweredBy},
		)
	}
	return out
}

func ConvertProjectsToProjectsOutput(source []database.Project) output.ProjectsOutput {
	out := output.ProjectsOutput{}

	if len(source) > 0 {
		for _, value := range source {
			out.Body = append(out.Body, output.Project{ID: value.ID, Name: value.Name, Description: value.Description})
		}
	} else {
		out.Body = make([]output.Project, 0)
	}

	return out
}

func ConvertRolesToRolesOutput(source []database.Role) output.RolesOutput {
	out := output.RolesOutput{}

	if len(source) > 0 {
		for _, value := range source {
			out.Body = append(out.Body, output.Role{ID: value.ID, Name: value.Name, Description: value.Description})
		}
	} else {
		out.Body = make([]output.Role, 0)
	}

	return out
}
