package output

type Role struct {
	ID          int64  `json:"id" example:"1" doc:"ID of the role"`
	Name        string `json:"name" example:"System admin" doc:"Role name"`
	Description string `json:"description" example:"This is a system admin" doc:"Role description"`
}

type UserWithRoleInfo struct {
	ID    int64  `json:"id" example:"1" doc:"ID of the user"`
	Email string `json:"email" format:"email" doc:"User email"`
	Role  string `json:"role" doc:"User role"`
}

type GetUserRoleOutput struct {
	Body UserWithRoleInfo
}

type RoleOutput struct {
	Body Role
}

type RolesOutput struct {
	Body []Role
}
