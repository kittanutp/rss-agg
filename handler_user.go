package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kittanutp/rss-agg/internal/database"
)

func (cfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}

	params := parameter{}
	decodeJSON(w, r, &params)

	user, user_err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
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

	respondWithJSON(w, 201, convertUserResponse(user))
}

func (cfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, convertUserResponse(user))
}
