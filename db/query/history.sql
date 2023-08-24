-- name: GetHistory :one
SELECT * FROM histories
WHERE id = $1 LIMIT 1;

-- name: ListHistories :many
SELECT * FROM histories
WHERE user_agent = $1 OR user_id = $2
ORDER BY last_read_at DESC
LIMIT $3
OFFSET $4;

-- name: ListHistoriesByAgent :many
SELECT * FROM histories
WHERE user_agent = $1
ORDER BY last_read_at DESC
LIMIT $2
OFFSET $3;

-- name: ListHistoriesByUserId :many
SELECT * FROM histories
WHERE user_id = $1
ORDER BY last_read_at DESC
LIMIT $2
OFFSET $3;

-- name: CreateHistory :one
INSERT INTO histories (
  user_agent,
  client_ip,
  user_id,
  mangadex_id,
  cover_image,
  title,
  reading_chapter,
  last_read_at,
  path,
  al_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateHistory :one
UPDATE histories
SET
 reading_chapter = $3,
 last_read_at = $4
WHERE mangadex_id = $1 AND user_agent = $2
RETURNING *;

-- name: UpdateHistoriesToUser :exec
UPDATE histories
SET
 user_id = $2
WHERE user_agent = $1 AND user_id IS NULL;

-- name: DeleteHistory :exec
DELETE FROM histories
WHERE id = $1;
