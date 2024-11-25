-- name: CreateProject :one
INSERT INTO projects (name, description) VALUES ($1, $2) RETURNING *;

-- name: GetProject :one
SELECT * FROM projects WHERE id = $1;

-- name: ListProjects :many
SELECT * FROM projects;

-- name: UpdateProject :one
UPDATE projects
  SET name = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects WHERE id = $1;