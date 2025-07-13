package main

import (
	"article-test/config"
	"article-test/internal/article/delivery"
	"article-test/internal/article/repository"
	"article-test/internal/article/usecase"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()
	if db == nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("Connected to database")

	cache := config.NewRedisClient().Connect()
	if cache == nil {
		log.Fatal("Failed to connect to Redis")
	}

	log.Println("Connected to Redis at", cache.Options().Addr)

	defer cache.Close()

	defer db.Close()

	if err := goose.Up(db, "./migration/db"); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewArticleRepository(db)
	service := usecase.NewArticleService(repo, cache)
	handler := delivery.NewArticleHandler(service)

	http.Handle("/articles", handler)
	log.Println("Server running at :8081")
	http.ListenAndServe(":8081", nil)
}
