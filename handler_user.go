package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kittanutp/rss-agg/internal/auth"
	"github.com/kittanutp/rss-agg/internal/database"
)

func (apiCfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Unable to decode json as:", err)
		respondWithError(w, 400, fmt.Sprintf("Unable to decode json as: %v", err))
		return
	}

	user, user_err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if user_err != nil {
		log.Println("Unable to create user as:", user_err)
		respondWithError(w, 400, fmt.Sprintf("Unable to create user as: %v", user_err))
		return
	}

	respondWithJSON(w, 200, convertUserResponse(user))
}

func (apiCfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		log.Println("Unable to get API key as:", err)
		respondWithError(w, 403, "Unauthorized")
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		log.Println("Unable to get user as:", err)
		respondWithError(w, 403, "Unauthorized")
		return
	}

	respondWithJSON(w, 200, convertUserResponse(user))
}
