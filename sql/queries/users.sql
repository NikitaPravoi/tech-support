-- name: CreateUser :one
INSERT INTO users (login, email, password, role_id)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserCredentials :one
SELECT id, login, password FROM users WHERE login = $1;

-- name: UpdateUser :exec
UPDATE users
SET login = $2, email = $3, password = $4
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetProjectUsers :many
SELECT
    users.id, users.login, email, r.name AS role
FROM users
         JOIN public.roles r on r.id = users.role_id
         JOIN public.user_projects up on users.id = up.user_id
WHERE project_id = $1;

-- name: GetUserWithHisRoles :one
SELECT u.id, u.email, r.name AS role
FROM users u
    JOIN roles r ON u.role_id = r.id
WHERE u.id = $1;

