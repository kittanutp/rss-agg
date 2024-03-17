// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: follows.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow WHERE feed_id = $1 and user_id = $2
`

type DeleteFeedFollowParams struct {
	FeedID uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.FeedID, arg.UserID)
	return err
}

const folows = `-- name: Folows :one
    INSERT INTO feed_follow (user_id, feed_id, created_at)
    VALUES ($1, $2, $3)
    RETURNING feed_id, user_id, created_at
`

type FolowsParams struct {
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
}

func (q *Queries) Folows(ctx context.Context, arg FolowsParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, folows, arg.UserID, arg.FeedID, arg.CreatedAt)
	var i FeedFollow
	err := row.Scan(&i.FeedID, &i.UserID, &i.CreatedAt)
	return i, err
}

const getFollowFeeds = `-- name: GetFollowFeeds :many
    SELECT feeds.id, feeds.user_id,  feeds.name, feeds.url, feeds.created_at, feeds.updated_at, feeds.last_fetch_at
    FROM feeds
    JOIN feed_follow ON feed_follow.feed_id = feeds.id
    WHERE feed_follow.user_id = $1
`

type GetFollowFeedsRow struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Name        string
	Url         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LastFetchAt sql.NullTime
}

func (q *Queries) GetFollowFeeds(ctx context.Context, userID uuid.UUID) ([]GetFollowFeedsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFollowFeeds, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFollowFeedsRow
	for rows.Next() {
		var i GetFollowFeedsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LastFetchAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
