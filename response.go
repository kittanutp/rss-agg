package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/kittanutp/rss-agg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func convertUserResponse(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
	}
}

type Feed struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Name        string     `json:"name"`
	Url         string     `json:"url"`
	UserID      uuid.UUID  `json:"user_id"`
	LastFetchAt *time.Time `json:"last_fetch_at"`
}

func convertFeedResponse(feed database.Feed) Feed {
	return Feed{
		ID:          feed.ID,
		CreatedAt:   feed.CreatedAt,
		UpdatedAt:   feed.UpdatedAt,
		Name:        feed.Name,
		Url:         feed.Url,
		UserID:      feed.UserID,
		LastFetchAt: nullTimeToTimePtr(feed.LastFetchAt),
	}
}

type Follow struct {
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func convertFollowResponse(folow database.FeedFollow) Follow {
	return Follow{
		FeedID:    folow.FeedID,
		UserID:    folow.UserID,
		CreatedAt: folow.CreatedAt,
	}
}

type FollowFeed struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Name        string     `json:"name"`
	Url         string     `json:"url"`
	UserID      uuid.UUID  `json:"user_id"`
	LastFetchAt *time.Time `json:"last_fetch_at"`
}

func convertFollowFeedResponse(feed database.GetFollowFeedsRow) Feed {
	return Feed{
		ID:          feed.ID,
		CreatedAt:   feed.CreatedAt,
		UpdatedAt:   feed.UpdatedAt,
		Name:        feed.Name,
		Url:         feed.Url,
		UserID:      feed.UserID,
		LastFetchAt: nullTimeToTimePtr(feed.LastFetchAt),
	}
}

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func convertPostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: nullStringToStringPtr(post.Description),
		PublishedAt: nullTimeToTimePtr(post.PublishedAt),
		FeedID:      post.FeedID,
	}
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}
