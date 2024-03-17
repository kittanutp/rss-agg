// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Url         string
	UserID      uuid.UUID
	LastFetchAt sql.NullTime
}

type FeedFollow struct {
	FeedID    uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}
