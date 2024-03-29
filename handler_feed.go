package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kittanutp/rss-agg/internal/database"
)

func (cfg *apiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := parameter{}
	decodeJSON(w, r, &params)

	feed, feed_err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if feed_err != nil {
		log.Println("Unable to create feed as:", feed_err)
		respondWithError(w, 400, fmt.Sprintf("Unable to create feed as: %v", feed_err))
		return
	}

	respondWithJSON(w, 200, convertFeedResponse(feed))
}

func (cfg *apiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, feed_err := cfg.DB.GetFeeds(r.Context())
	if feed_err != nil {
		log.Println("Unable to get feeds as:", feed_err)
		respondWithError(w, 400, fmt.Sprintf("Unable to get feeds as: %v", feed_err))
		return
	}
	var resp_feeds []Feed
	for _, feed := range feeds {
		resp_feeds = append(resp_feeds, convertFeedResponse(feed))
	}

	respondWithJSON(w, 200, resp_feeds)
}
