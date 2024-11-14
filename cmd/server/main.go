package main

import (
	"detour/internal/application/shortener"
	"detour/internal/domain/url"
	"detour/internal/infrastructure/http"
	"detour/internal/infrastructure/http/handlers"
	"detour/internal/infrastructure/mysql"
	"detour/internal/infrastructure/mysql/repositories"
	"log"
	"time"
)

func main() {
	// Initialize repository
	mysqlClient, err := mysql.NewClient(mysql.Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "",
		Database: "shortcut",
	})
	if err != nil {
		log.Fatalf("Failed to initialize MySQL client: %v", err)
	}

	repo := repositories.NewURLRepo(mysqlClient.DB())
	service := url.NewService(repo)
	shortenerUseCase := shortener.NewUseCase(service)

	urlHandler := handlers.NewURLHandler(shortenerUseCase)

	router := http.Setup(urlHandler)

	server := http.NewServer(http.ServerConfig{
		Port:           "8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		HandlerTimeout: 10 * time.Second,
	}, router)

	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
