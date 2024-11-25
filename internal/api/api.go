package api

import (
	"github.com/danielgtaylor/huma/v2"
	"net/http"
	"tech-support/internal/api/constants"
	"tech-support/internal/handlers"
)

var operationSecurity []map[string][]string

func RegisterRoutes(api huma.API) {
	operationSecurity = []map[string][]string{{"auth": {"auth"}}}

	registerAuthorization(api)
	registerSupportTickets(api)
	registerProjects(api)
	registerRoles(api)
	registerUserProjects(api)
}

func registerAuthorization(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "login",
		Method:        http.MethodPost,
		Path:          "/login",
		Summary:       "Log in",
		Description:   "Logging in",
		Tags:          []string{constants.Authorization},
		DefaultStatus: http.StatusOK,
	}, handlers.LoginHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "logout",
		Method:        http.MethodGet,
		Path:          "/logout",
		Summary:       "Log out",
		Description:   "Logging out",
		Tags:          []string{constants.Authorization},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.LogoutHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "register",
		Method:        http.MethodPost,
		Path:          "/register",
		Summary:       "Register",
		Description:   "Register a user",
		Tags:          []string{constants.Authorization},
		DefaultStatus: http.StatusCreated,
	}, handlers.CreateUserHandler)
}

func registerSupportTickets(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "create-ticket",
		Method:        http.MethodPost,
		Path:          "/ticket",
		Summary:       "Create ticket",
		Description:   "Create support ticket with a starting status with 0",
		Tags:          []string{constants.Support},
		DefaultStatus: http.StatusCreated,
		Security:      operationSecurity,
	}, handlers.CreateTicketHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "get-ticket",
		Method:        http.MethodGet,
		Path:          "/ticket/{id}",
		Summary:       "Get ticket",
		Description:   "Get support ticket",
		Tags:          []string{constants.Support},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.GetTicketHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "list-tickets",
		Method:        http.MethodGet,
		Path:          "/ticket",
		Summary:       "List tickets",
		Description:   "Get support tickets",
		Tags:          []string{constants.Support},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.ListTicketHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "list-tickets",
		Method:        http.MethodDelete,
		Path:          "/ticket/{id}",
		Summary:       "Delete ticket",
		Description:   "Delete support ticket",
		Tags:          []string{constants.Support},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.DeleteTicketHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "update-ticket",
		Method:        http.MethodPut,
		Path:          "/ticket/{id}",
		Summary:       "Update ticket",
		Description:   "Update support ticket",
		Tags:          []string{constants.Support},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.UpdateTicketHandler)
}

func registerProjects(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "create-project",
		Method:        http.MethodPost,
		Path:          "/admin/project",
		Summary:       "Create project",
		Description:   "Create project that requires support tickets",
		Tags:          []string{constants.Project},
		DefaultStatus: http.StatusCreated,
		Security:      operationSecurity,
	}, handlers.CreateProjectHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "get-project",
		Method:        http.MethodGet,
		Path:          "/project/{id}",
		Summary:       "Get project",
		Description:   "Get support project",
		Tags:          []string{constants.Project},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.GetProjectHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "list-projects",
		Method:        http.MethodGet,
		Path:          "/project",
		Summary:       "List project",
		Description:   "Get support project",
		Tags:          []string{constants.Project},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.ListProjectHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "delete-project",
		Method:        http.MethodDelete,
		Path:          "/admin/project/{id}",
		Summary:       "Delete project",
		Description:   "Delete support project",
		Tags:          []string{constants.Project},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.DeleteProjectHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "update-project",
		Method:        http.MethodPut,
		Path:          "/admin/project/{id}",
		Summary:       "Update project",
		Description:   "Update support project",
		Tags:          []string{constants.Project},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.UpdateProjectHandler)
}

func registerRoles(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "create-role",
		Method:        http.MethodPost,
		Path:          "/admin/role",
		Summary:       "Create role",
		Description:   "Create new role",
		Tags:          []string{constants.Roles},
		DefaultStatus: http.StatusCreated,
		Security:      operationSecurity,
	}, handlers.CreateRoleHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "get-user-role",
		Method:        http.MethodGet,
		Path:          "/role",
		Summary:       "Get user role",
		Description:   "Get user support role",
		Tags:          []string{constants.User, constants.Roles},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.GetUserRoleHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "get-role",
		Method:        http.MethodGet,
		Path:          "/role/{id}",
		Summary:       "Get role",
		Description:   "Get support role",
		Tags:          []string{constants.Roles},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.GetRoleHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "list-Roles",
		Method:        http.MethodGet,
		Path:          "/admin/role",
		Summary:       "List role",
		Description:   "Get support role",
		Tags:          []string{constants.Roles},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.ListRoleHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "list-role",
		Method:        http.MethodDelete,
		Path:          "/admin/role/{id}",
		Summary:       "Delete role",
		Description:   "Delete support role",
		Tags:          []string{constants.Roles},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.DeleteRoleHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "update-role",
		Method:        http.MethodPut,
		Path:          "/admin/role/{id}",
		Summary:       "Update role",
		Description:   "Update support role",
		Tags:          []string{constants.Roles},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.UpdateRoleHandler)
}

func registerUserProjects(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "create-user-project",
		Method:        http.MethodPost,
		Path:          "/admin/user-project",
		Summary:       "Create user project",
		Description:   "Create user project binding",
		Tags:          []string{constants.User, constants.Project},
		DefaultStatus: http.StatusCreated,
		Security:      operationSecurity,
	}, handlers.CreateUserProjectHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "list-user-project",
		Method:        http.MethodGet,
		Path:          "/user-project",
		Summary:       "List user project",
		Description:   "Get all user project bindings",
		Tags:          []string{constants.User, constants.Project},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.ListUserProjectHandler)
	huma.Register(api, huma.Operation{
		OperationID:   "delete-user-project",
		Method:        http.MethodDelete,
		Path:          "/admin/user-project/{id}",
		Summary:       "Delete user project",
		Description:   "Delete user project binding",
		Tags:          []string{constants.User, constants.Project},
		DefaultStatus: http.StatusOK,
		Security:      operationSecurity,
	}, handlers.DeleteUserProjectHandler)
}
