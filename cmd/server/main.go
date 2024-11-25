package main

import (
	"database/sql"
	"errors"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"tech-support/internal/api"
	"tech-support/internal/api/constants"
	"tech-support/internal/database"
	"tech-support/internal/services"
)

var humaAPI huma.API

// SessionMiddleware is used to perform session validation and authentication support
func SessionMiddleware(ctx huma.Context, next func(huma.Context)) {
	if len(ctx.Operation().Security) == 0 {
		next(ctx)
		return
	}

	token := strings.TrimPrefix(ctx.Header("Authorization"), "Bearer ")
	if len(token) == 0 {
		huma.WriteErr(humaAPI, ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	session := services.CheckTokenValid(ctx.Context(), token)
	if session == nil {
		huma.WriteErr(humaAPI, ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Used to define what user is authenticated
	ctx = huma.WithValue(ctx, "user_id", session.UserID)

	// Validate user role after user_id was assigned
	RolesMiddleware(ctx, next)
}

// RolesMiddleware is used to identify user permissions on various operations
func RolesMiddleware(ctx huma.Context, next func(huma.Context)) {
	context := ctx.Context()

	user, err := database.Repository.GetUser(context, context.Value("user_id").(int64))
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			huma.WriteErr(humaAPI, ctx, http.StatusBadRequest, "User does not exist")
			return
		}
		huma.WriteErr(humaAPI, ctx, http.StatusInternalServerError, "Database error")
		return
	}

	url := ctx.URL()
	if strings.Contains(url.String(), "admin") {
		if constants.Role(user.RoleID) != constants.SystemAdmin && constants.Role(user.RoleID) != constants.TechnicalSupport {
			huma.WriteErr(humaAPI, ctx, http.StatusForbidden, "Access denied")
			return
		}
	}

	next(ctx)
}

func main() {
	db := database.MustInitDB()
	defer db.Close()

	router := chi.NewMux()
	config := huma.DefaultConfig("tech-support", "1.0.0")
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"auth": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
			Flows: &huma.OAuthFlows{
				AuthorizationCode: &huma.OAuthFlow{
					AuthorizationURL: "https://example.com/login",
					Scopes: map[string]string{
						"auth": "basic auth",
					},
				},
			},
		},
	}

	humaAPI = humachi.New(router, config)
	humaAPI.UseMiddleware(SessionMiddleware)

	api.RegisterRoutes(humaAPI)

	// Start the server!
	http.ListenAndServe(":8000", router)
}
