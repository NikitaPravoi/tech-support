-- name: CreateUserProject :one
INSERT INTO user_projects (user_id, project_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserProject :one
SELECT * FROM user_projects WHERE id = $1;

-- name: ListUserProjects :many
SELECT p.id, p.name, p.description FROM user_projects
JOIN public.projects p on user_projects.project_id = p.id
WHERE user_id = $1;

-- name: UpdateUserProject :one
UPDATE user_projects
SET user_id = $2, project_id = $3
WHERE id = $1
RETURNING *;


-- name: DeleteUserProject :exec
DELETE FROM user_projects WHERE user_id = $1 AND project_id = $2;