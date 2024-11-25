package input

type CreateProjectInput struct {
	Body struct {
		Name        string  `json:"name" example:"Project" doc:"Project name"`
		Description *string `json:"description" example:"My favourite project" doc:"Project description"`
	}
}

type UpdateProjectInput struct {
	ID   int64 `json:"id" example:"1" doc:"ID" path:"id"`
	Body struct {
		Name        string  `json:"name" example:"Project" doc:"Project name"`
		Description *string `json:"description" example:"My favourite project" doc:"Project description"`
	}
}
