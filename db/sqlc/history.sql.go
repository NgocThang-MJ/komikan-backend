// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: history.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createHistory = `-- name: CreateHistory :one
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
RETURNING id, user_agent, client_ip, user_id, al_id, mangadex_id, cover_image, title, reading_chapter, path, last_read_at, created_at
`

type CreateHistoryParams struct {
	UserAgent      string      `json:"user_agent"`
	ClientIp       string      `json:"client_ip"`
	UserID         pgtype.UUID `json:"user_id"`
	MangadexID     string      `json:"mangadex_id"`
	CoverImage     string      `json:"cover_image"`
	Title          string      `json:"title"`
	ReadingChapter string      `json:"reading_chapter"`
	LastReadAt     time.Time   `json:"last_read_at"`
	Path           string      `json:"path"`
	AlID           string      `json:"al_id"`
}

func (q *Queries) CreateHistory(ctx context.Context, arg CreateHistoryParams) (History, error) {
	row := q.db.QueryRow(ctx, createHistory,
		arg.UserAgent,
		arg.ClientIp,
		arg.UserID,
		arg.MangadexID,
		arg.CoverImage,
		arg.Title,
		arg.ReadingChapter,
		arg.LastReadAt,
		arg.Path,
		arg.AlID,
	)
	var i History
	err := row.Scan(
		&i.ID,
		&i.UserAgent,
		&i.ClientIp,
		&i.UserID,
		&i.AlID,
		&i.MangadexID,
		&i.CoverImage,
		&i.Title,
		&i.ReadingChapter,
		&i.Path,
		&i.LastReadAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteHistory = `-- name: DeleteHistory :exec
DELETE FROM histories
WHERE id = $1
`

func (q *Queries) DeleteHistory(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteHistory, id)
	return err
}

const getHistory = `-- name: GetHistory :one
SELECT id, user_agent, client_ip, user_id, al_id, mangadex_id, cover_image, title, reading_chapter, path, last_read_at, created_at FROM histories
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetHistory(ctx context.Context, id pgtype.UUID) (History, error) {
	row := q.db.QueryRow(ctx, getHistory, id)
	var i History
	err := row.Scan(
		&i.ID,
		&i.UserAgent,
		&i.ClientIp,
		&i.UserID,
		&i.AlID,
		&i.MangadexID,
		&i.CoverImage,
		&i.Title,
		&i.ReadingChapter,
		&i.Path,
		&i.LastReadAt,
		&i.CreatedAt,
	)
	return i, err
}

const listHistories = `-- name: ListHistories :many
SELECT id, user_agent, client_ip, user_id, al_id, mangadex_id, cover_image, title, reading_chapter, path, last_read_at, created_at FROM histories
WHERE user_agent = $1 OR user_id = $2
ORDER BY last_read_at DESC
LIMIT $3
OFFSET $4
`

type ListHistoriesParams struct {
	UserAgent string      `json:"user_agent"`
	UserID    pgtype.UUID `json:"user_id"`
	Limit     int32       `json:"limit"`
	Offset    int32       `json:"offset"`
}

func (q *Queries) ListHistories(ctx context.Context, arg ListHistoriesParams) ([]History, error) {
	rows, err := q.db.Query(ctx, listHistories,
		arg.UserAgent,
		arg.UserID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []History{}
	for rows.Next() {
		var i History
		if err := rows.Scan(
			&i.ID,
			&i.UserAgent,
			&i.ClientIp,
			&i.UserID,
			&i.AlID,
			&i.MangadexID,
			&i.CoverImage,
			&i.Title,
			&i.ReadingChapter,
			&i.Path,
			&i.LastReadAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listHistoriesByAgent = `-- name: ListHistoriesByAgent :many
SELECT id, user_agent, client_ip, user_id, al_id, mangadex_id, cover_image, title, reading_chapter, path, last_read_at, created_at FROM histories
WHERE user_agent = $1
ORDER BY last_read_at DESC
LIMIT $2
OFFSET $3
`

type ListHistoriesByAgentParams struct {
	UserAgent string `json:"user_agent"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

func (q *Queries) ListHistoriesByAgent(ctx context.Context, arg ListHistoriesByAgentParams) ([]History, error) {
	rows, err := q.db.Query(ctx, listHistoriesByAgent, arg.UserAgent, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []History{}
	for rows.Next() {
		var i History
		if err := rows.Scan(
			&i.ID,
			&i.UserAgent,
			&i.ClientIp,
			&i.UserID,
			&i.AlID,
			&i.MangadexID,
			&i.CoverImage,
			&i.Title,
			&i.ReadingChapter,
			&i.Path,
			&i.LastReadAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listHistoriesByUserId = `-- name: ListHistoriesByUserId :many
SELECT id, user_agent, client_ip, user_id, al_id, mangadex_id, cover_image, title, reading_chapter, path, last_read_at, created_at FROM histories
WHERE user_id = $1
ORDER BY last_read_at DESC
LIMIT $2
OFFSET $3
`

type ListHistoriesByUserIdParams struct {
	UserID pgtype.UUID `json:"user_id"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

func (q *Queries) ListHistoriesByUserId(ctx context.Context, arg ListHistoriesByUserIdParams) ([]History, error) {
	rows, err := q.db.Query(ctx, listHistoriesByUserId, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []History{}
	for rows.Next() {
		var i History
		if err := rows.Scan(
			&i.ID,
			&i.UserAgent,
			&i.ClientIp,
			&i.UserID,
			&i.AlID,
			&i.MangadexID,
			&i.CoverImage,
			&i.Title,
			&i.ReadingChapter,
			&i.Path,
			&i.LastReadAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateHistoriesToUser = `-- name: UpdateHistoriesToUser :exec
UPDATE histories
SET
 user_id = $2
WHERE user_agent = $1 AND user_id IS NULL
`

type UpdateHistoriesToUserParams struct {
	UserAgent string      `json:"user_agent"`
	UserID    pgtype.UUID `json:"user_id"`
}

func (q *Queries) UpdateHistoriesToUser(ctx context.Context, arg UpdateHistoriesToUserParams) error {
	_, err := q.db.Exec(ctx, updateHistoriesToUser, arg.UserAgent, arg.UserID)
	return err
}

const updateHistory = `-- name: UpdateHistory :one
UPDATE histories
SET
 reading_chapter = $3,
 last_read_at = $4
WHERE mangadex_id = $1 AND user_agent = $2
RETURNING id, user_agent, client_ip, user_id, al_id, mangadex_id, cover_image, title, reading_chapter, path, last_read_at, created_at
`

type UpdateHistoryParams struct {
	MangadexID     string    `json:"mangadex_id"`
	UserAgent      string    `json:"user_agent"`
	ReadingChapter string    `json:"reading_chapter"`
	LastReadAt     time.Time `json:"last_read_at"`
}

func (q *Queries) UpdateHistory(ctx context.Context, arg UpdateHistoryParams) (History, error) {
	row := q.db.QueryRow(ctx, updateHistory,
		arg.MangadexID,
		arg.UserAgent,
		arg.ReadingChapter,
		arg.LastReadAt,
	)
	var i History
	err := row.Scan(
		&i.ID,
		&i.UserAgent,
		&i.ClientIp,
		&i.UserID,
		&i.AlID,
		&i.MangadexID,
		&i.CoverImage,
		&i.Title,
		&i.ReadingChapter,
		&i.Path,
		&i.LastReadAt,
		&i.CreatedAt,
	)
	return i, err
}