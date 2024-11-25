package output

type Project struct {
	ID          int64   `json:"id" example:"1" doc:""`
	Name        string  `json:"name" example:"Tech-support" doc:"Project name"`
	Description *string `json:"description" example:"This is my favourite project" doc:"Project description"`
}

type ProjectOutput struct {
	Body Project
}

type ProjectsOutput struct {
	Body []Project
}
