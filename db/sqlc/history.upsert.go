package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Insert params order matter! If we prevent the UserAgent conflict, the user_agent should come before the user_id and vice versa
const upsertHistoryAgent = `-- name: UpsertHistoryAgent :exec
INSERT INTO histories (
  user_agent,
  mangadex_id,
  client_ip,
  user_id,
  cover_image,
  title,
  reading_chapter,
  last_read_at,
  al_id,
  path
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
ON CONFLICT (mangadex_id,user_agent)
DO UPDATE SET
  user_agent = EXCLUDED.user_agent,
  reading_chapter = EXCLUDED.reading_chapter,
  path = EXCLUDED.path,
  cover_image = EXCLUDED.cover_image,
  last_read_at = EXCLUDED.last_read_at
`

const upsertHistoryUserId = `-- name: UpsertHistoryUserId :exec
INSERT INTO histories (
  user_id,
  mangadex_id,
  client_ip,
  user_agent,
  cover_image,
  title,
  reading_chapter,
  last_read_at,
  al_id,
  path
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
ON CONFLICT (mangadex_id,user_id)
DO UPDATE SET
  user_agent = EXCLUDED.user_agent,
  reading_chapter = EXCLUDED.reading_chapter,
  path = EXCLUDED.path,
  cover_image = EXCLUDED.cover_image,
  last_read_at = EXCLUDED.last_read_at
`

type UpsertHistoryParams struct {
	UserAgent      string      `json:"user_agent"`
	ClientIp       string      `json:"client_ip"`
	UserID         pgtype.UUID `json:"user_id"`
	MangadexID     string      `json:"mangadex_id"`
	CoverImage     string      `json:"cover_image"`
	Title          string      `json:"title"`
	ReadingChapter string      `json:"reading_chapter"`
	AlId           string      `json:"al_id"`
	Path           string      `json:"path"`
	LastReadAt     time.Time   `json:"last_read_at"`
}

func (q *Queries) UpsertHistoryAgent(
	ctx context.Context,
	arg UpsertHistoryParams,
) error {
	_, err := q.db.Exec(ctx, upsertHistoryAgent,
		arg.UserAgent,
		arg.MangadexID,
		arg.ClientIp,
		arg.UserID,
		arg.CoverImage,
		arg.Title,
		arg.ReadingChapter,
		arg.LastReadAt,
		arg.AlId,
		arg.Path,
	)
	return err
}

func (q *Queries) UpsertHistoryUserId(
	ctx context.Context,
	arg UpsertHistoryParams,
) error {
	_, err := q.db.Exec(ctx, upsertHistoryUserId,
		arg.UserID,
		arg.MangadexID,
		arg.ClientIp,
		arg.UserAgent,
		arg.CoverImage,
		arg.Title,
		arg.ReadingChapter,
		arg.LastReadAt,
		arg.AlId,
		arg.Path,
	)
	return err
}
