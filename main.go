package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/kittanutp/rss-agg/internal/database"

	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Start Up Server")
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	fmt.Printf("Listening on port %v \n", portString)

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found")
	}

	conn, db_err := sql.Open("postgres", dbURL)
	if db_err != nil {
		log.Fatal("Unable to connect database as :", db_err)
	}

	cfg := apiConfig{
		DB: database.New(conn),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Default().Handler)

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) { respondWithJSON(w, 200, "OK :)") })
	r.Get("/test-error", func(w http.ResponseWriter, r *http.Request) { respondWithError(w, 400, "test-error") })

	userRouter := chi.NewRouter()
	userRouter.Post("/new", cfg.HandlerCreateUser)
	userRouter.Get("/info", cfg.middlewareAuth(cfg.HandlerGetUser))

	feedRouter := chi.NewRouter()
	feedRouter.Post("/new", cfg.middlewareAuth(cfg.HandlerCreateFeed))
	feedRouter.Post("/follow", cfg.middlewareAuth(cfg.HandlerFollow))
	feedRouter.Get("/follow", cfg.middlewareAuth(cfg.HandlerGetFollowFeeds))
	feedRouter.Delete("/follow/{feedFollowID}", cfg.middlewareAuth(cfg.HandlerFeedFollowDelete))
	feedRouter.Get("/all", cfg.HandlerGetFeeds)

	postRouter := chi.NewRouter()
	postRouter.Get("/all", cfg.middlewareAuth(cfg.HandlerPostsGet))

	r.Mount("/user", userRouter)
	r.Mount("/feed", feedRouter)
	r.Mount("/post", postRouter)

	srv := &http.Server{
		Addr:           ":" + portString,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(cfg.DB, collectionConcurrency, collectionInterval)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
