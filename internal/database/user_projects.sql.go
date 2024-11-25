// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user_projects.sql

package database

import (
	"context"
)

const createUserProject = `-- name: CreateUserProject :one
INSERT INTO user_projects (user_id, project_id)
VALUES ($1, $2)
RETURNING id, user_id, project_id
`

type CreateUserProjectParams struct {
	UserID    int64 `json:"user_id"`
	ProjectID int64 `json:"project_id"`
}

func (q *Queries) CreateUserProject(ctx context.Context, arg CreateUserProjectParams) (UserProject, error) {
	row := q.db.QueryRow(ctx, createUserProject, arg.UserID, arg.ProjectID)
	var i UserProject
	err := row.Scan(&i.ID, &i.UserID, &i.ProjectID)
	return i, err
}

const deleteUserProject = `-- name: DeleteUserProject :exec
DELETE FROM user_projects WHERE user_id = $1 AND project_id = $2
`

type DeleteUserProjectParams struct {
	UserID    int64 `json:"user_id"`
	ProjectID int64 `json:"project_id"`
}

func (q *Queries) DeleteUserProject(ctx context.Context, arg DeleteUserProjectParams) error {
	_, err := q.db.Exec(ctx, deleteUserProject, arg.UserID, arg.ProjectID)
	return err
}

const getUserProject = `-- name: GetUserProject :one
SELECT id, user_id, project_id FROM user_projects WHERE id = $1
`

func (q *Queries) GetUserProject(ctx context.Context, id int64) (UserProject, error) {
	row := q.db.QueryRow(ctx, getUserProject, id)
	var i UserProject
	err := row.Scan(&i.ID, &i.UserID, &i.ProjectID)
	return i, err
}

const listUserProjects = `-- name: ListUserProjects :many
SELECT p.id, p.name, p.description FROM user_projects
JOIN public.projects p on user_projects.project_id = p.id
WHERE user_id = $1
`

func (q *Queries) ListUserProjects(ctx context.Context, userID int64) ([]Project, error) {
	rows, err := q.db.Query(ctx, listUserProjects, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserProject = `-- name: UpdateUserProject :one
UPDATE user_projects
SET user_id = $2, project_id = $3
WHERE id = $1
RETURNING id, user_id, project_id
`

type UpdateUserProjectParams struct {
	ID        int64 `json:"id"`
	UserID    int64 `json:"user_id"`
	ProjectID int64 `json:"project_id"`
}

func (q *Queries) UpdateUserProject(ctx context.Context, arg UpdateUserProjectParams) (UserProject, error) {
	row := q.db.QueryRow(ctx, updateUserProject, arg.ID, arg.UserID, arg.ProjectID)
	var i UserProject
	err := row.Scan(&i.ID, &i.UserID, &i.ProjectID)
	return i, err
}
