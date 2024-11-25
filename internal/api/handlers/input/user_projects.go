package input

type CreateUserProjectInput struct {
	Body struct {
		UserID    int64 `json:"user_id" example:"1" doc:"ID of the user"`
		ProjectID int64 `json:"project_id" example:"1" doc:"ID of the project"`
	}
}

type UpdateUserProjectInput struct {
	ID   int64 `json:"id" example:"1" doc:"ID" path:"id"`
	Body struct {
		UserID    int64 `json:"user_id" example:"1" doc:"ID of the user"`
		ProjectID int64 `json:"project_id" example:"1" doc:"ID of the project"`
	}
}
