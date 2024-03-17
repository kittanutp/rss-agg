package main

import (
	"net/http"
	"strconv"

	"github.com/kittanutp/rss-agg/internal/database"
)

func (cfg *apiConfig) HandlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := cfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get posts for user")
		return
	}

	var resp_posts []Post
	for _, post := range posts {
		resp_posts = append(resp_posts, convertPostToPost(post))
	}

	respondWithJSON(w, http.StatusOK, resp_posts)
}
