// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: projects.sql

package database

import (
	"context"
)

const createProject = `-- name: CreateProject :one
INSERT INTO projects (name, description) VALUES ($1, $2) RETURNING id, name, description
`

type CreateProjectParams struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, createProject, arg.Name, arg.Description)
	var i Project
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM projects WHERE id = $1
`

func (q *Queries) DeleteProject(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProject, id)
	return err
}

const getProject = `-- name: GetProject :one
SELECT id, name, description FROM projects WHERE id = $1
`

func (q *Queries) GetProject(ctx context.Context, id int64) (Project, error) {
	row := q.db.QueryRow(ctx, getProject, id)
	var i Project
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const listProjects = `-- name: ListProjects :many
SELECT id, name, description FROM projects
`

func (q *Queries) ListProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.Query(ctx, listProjects)
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

const updateProject = `-- name: UpdateProject :one
UPDATE projects
  SET name = $2, description = $3
WHERE id = $1
RETURNING id, name, description
`

type UpdateProjectParams struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, updateProject, arg.ID, arg.Name, arg.Description)
	var i Project
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}