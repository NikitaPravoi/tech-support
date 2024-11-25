package output

import "time"

type SupportTicket struct {
	ID          int64     `json:"id" example:"1" doc:"ID of the support ticket"`
	ProjectID   int64     `json:"project_id" example:"1" doc:"ID of the project"`
	StatusID    int64     `json:"status_id" example:"1" doc:"ID of the status of the support ticket"`
	Title       string    `json:"title" example:"Stuck in a washing machine" doc:"Title of the support ticket"`
	Description *string   `json:"description" example:"Step-bro help me, i'm stuck!" doc:"Description of the support ticket"`
	CreatedAt   time.Time `json:"created_at" format:"date-time" doc:"Created time of the support ticket"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time" doc:"Updated time of the support ticket"`
	CreatedBy   string    `json:"created_by" example:"Dispatcher" doc:"Creator of the support ticket"`
	Answer      *string   `json:"answer" example:"You are an idiot." doc:"Answer of the support ticket"`
	AnsweredBy  *int64    `json:"answered_by" example:"1" doc:"Who answered the support ticket"`
}

type SupportTicketOutput struct {
	Body SupportTicket
}

type SupportTicketsOutput struct {
	Body []SupportTicket
}
