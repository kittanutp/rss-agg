package main

import (
	"log"
	"net/http"

	"github.com/kittanutp/rss-agg/internal/auth"
	"github.com/kittanutp/rss-agg/internal/database"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			log.Println("Unable to get API key as:", err)
			respondWithError(w, 403, "Unauthorized")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			log.Println("Unable to get user as:", err)
			respondWithError(w, 403, "Unauthorized")
			return
		}

		handler(w, r, user)
	}
}
