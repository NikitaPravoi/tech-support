package input

type CreateRoleInput struct {
	Body struct {
		Name        string `json:"name" example:"Project" doc:"Role name"`
		Description string `json:"description" example:"Admin" doc:"Role description"`
	}
}

type UpdateRoleInput struct {
	ID   int64 `json:"id" example:"1" doc:"ID" path:"id"`
	Body struct {
		Name        string `json:"name" example:"Project" doc:"Role name"`
		Description string `json:"description" example:"Admin" doc:"Role description"`
	}
}
