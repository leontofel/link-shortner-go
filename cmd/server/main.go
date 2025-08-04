package main

import (
	"context"
	"log"
	"net/http"

	"github.com/leontofel/link-shortener/internal/repository"
	"github.com/leontofel/link-shortener/internal/service"
	"github.com/leontofel/link-shortener/internal/handler"
)

func main() {
	ctx := context.Background()
	repo := repository.NewRedisRepository(ctx, "redis:6379")
	svc := service.NewShortenerService(repo)
	h := handler.NewHTTPHandler(svc)

	http.HandleFunc("/shorten", h.ShortenURL)
	http.HandleFunc("/", h.ResolveURL)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
