package input

type CreateSupportTicketInput struct {
	Body struct {
		ProjectID   int64   `json:"project_id" example:"1" doc:"Project to assign to"`
		Title       string  `json:"title" example:"Authorization" doc:"Title of a ticket"`
		Description *string `json:"description" example:"some kind of long message" doc:"Description of a ticket"`
		CreatedBy   string  `json:"created_by" example:"Dispatcher Chelyabinsk" doc:"Tickets author"`
	}
}

type UpdateSupportTicketInput struct {
	ID   int64 `json:"id" example:"1" doc:"ID of support ticket" path:"id"`
	Body struct {
		ProjectID   int64   `json:"project_id" example:"1" doc:"Project ID"`
		StatusID    int64   `json:"status_id" example:"1" doc:" Status ID"`
		Title       string  `json:"title" example:"Sample Title" doc:"Title of a ticket"`
		Description *string `json:"description" example:"Sample description" doc:"Description of a ticket"`
		CreatedBy   string  `json:"created_by" format:"date-time" doc:""`
	}
}
