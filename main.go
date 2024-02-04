package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Start Up Server")
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	fmt.Printf("Listening on port %v \n", portString)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Default().Handler)

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) { respondWithJSON(w, 200, "OK :)") })
	r.Get("/test-error", func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, 400, "test-error")
	})
	V1Router := chi.NewRouter()
	r.Mount("/v1", V1Router)

	srv := &http.Server{
		Addr:           ":" + portString,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
