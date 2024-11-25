-- name: CreateSession :one
INSERT INTO sessions (user_id, token)
VALUES ($1, $2) RETURNING *;

-- name: ListSessions :many
SELECT * FROM sessions;

-- name: GetSession :one
SELECT * FROM sessions WHERE id = $1;

-- name: GetSessionByToken :one
SELECT * FROM sessions WHERE token = $1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1;

-- name: DeleteSessionsByUserID :exec
DELETE FROM sessions
WHERE user_id = $1;

-- name: UpdateSessionExpiry :exec
UPDATE sessions
SET expires_at = CURRENT_TIMESTAMP
WHERE token = $1;