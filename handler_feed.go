package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kittanutp/rss-agg/internal/database"
)

func (apiCfg *apiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Unable to decode json as:", err)
		respondWithError(w, 400, fmt.Sprintf("Unable to decode json as: %v", err))
		return
	}

	feed, feed_err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
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

func (apiCfg *apiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	// type parameter struct {
	// 	Name string `json:"name"`
	// 	Url  string `json:"url"`
	// }
	// decoder := json.NewDecoder(r.Body)

	// params := parameter{}
	// err := decoder.Decode(&params)
	// if err != nil {
	// 	log.Println("Unable to decode json as:", err)
	// 	respondWithError(w, 400, fmt.Sprintf("Unable to decode json as: %v", err))
	// 	return
	// }

	feeds, feed_err := apiCfg.DB.GetFeeds(r.Context())
	if feed_err != nil {
		log.Println("Unable to create feed as:", feed_err)
		respondWithError(w, 400, fmt.Sprintf("Unable to get feeds as: %v", feed_err))
		return
	}
	var resp_feeds []Feed
	for _, feed := range feeds {
		resp_feeds = append(resp_feeds, convertFeedResponse(feed))
	}

	respondWithJSON(w, 200, resp_feeds)
}
