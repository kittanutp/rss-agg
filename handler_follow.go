package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/kittanutp/rss-agg/internal/database"
)

func (apiCfg *apiConfig) HandlerFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameter{}
	decodeJSON(w, r, &params)
	folow, err := apiCfg.DB.Folows(r.Context(), database.FolowsParams{
		CreatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		log.Println("Unable to folow feed as:", err)
		respondWithError(w, 400, fmt.Sprintf("Unable to folow feed as: %v", err))
		return
	}

	respondWithJSON(w, 200, convertFollowResponse(folow))
}

func (apiCfg *apiConfig) HandlerGetFollowFeeds(w http.ResponseWriter, r *http.Request, user database.User) {

	feeds, err := apiCfg.DB.GetFollowFeeds(r.Context(), user.ID)
	if err != nil {
		log.Println("Unable to create feed as:", err)
		respondWithError(w, 400, fmt.Sprintf("Unable to get feeds as: %v", err))
		return
	}
	var resp_feeds []Feed
	for _, feed := range feeds {
		resp_feeds = append(resp_feeds, convertFollowFeedResponse(feed))
	}

	respondWithJSON(w, 200, resp_feeds)
}

func (cfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feedFollowID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to unfollow feed")
		return
	}

	respondWithJSON(w, http.StatusOK, struct {
		Resp string `json:"resp"`
	}{
		Resp: "OK",
	})
}
