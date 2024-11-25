-- name: CreateSupportTicket :one
INSERT INTO support_tickets (project_id, title, description, created_by)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetSupportTicket :one
SELECT * FROM support_tickets WHERE id = $1;

-- name: ListSupportTickets :many
SELECT * FROM support_tickets;

-- name: ListSupportTicketsByProject :many
SELECT * FROM support_tickets WHERE project_id = $1;

-- name: UpdateSupportTicket :one
UPDATE support_tickets
SET project_id = $2, status_id = $3, title = $4, description = $5, updated_at = now(), created_by = $6
WHERE id = $1
RETURNING *;

-- name: DeleteSupportTicket :exec
DELETE FROM support_tickets WHERE id = $1;

-- name: WriteAnswerSupportTicket :one
UPDATE support_tickets
SET answer = $2, answered_by = $3
WHERE id = $1
RETURNING *;