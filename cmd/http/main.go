package main

import (
	"time"

	"github.com/BasbustDama/tiny-url/internal/generator"
	"github.com/BasbustDama/tiny-url/internal/handler"
	"github.com/BasbustDama/tiny-url/internal/storage"
	"github.com/BasbustDama/tiny-url/internal/usecase"
	"github.com/BasbustDama/tiny-url/pkg/cache"
	"github.com/BasbustDama/tiny-url/pkg/server"
)

const (
	defaultGenLength = 6
)

func main() {
	cache := cache.New(time.Hour, time.Hour)
	storage := storage.New(cache)

	generator := generator.New(defaultGenLength)

	usecase := usecase.New(storage, generator)

	handler := handler.NewHandler(usecase)

	server := server.New(handler, ":8080")
	server.Run()
}
