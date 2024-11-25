package output

type UserProject struct {
	ID        int64 `json:"id" example:"1" doc:"ID of the user project binding"`
	UserID    int64 `json:"user_id" example:"1" doc:"ID of the user"`
	ProjectID int64 `json:"project_id" example:"1" doc:"ID of the project"`
}

type UserProjectOutput struct {
	Body UserProject
}

type UserProjectsOutput struct {
	Body []UserProject
}
