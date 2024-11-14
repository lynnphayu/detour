package main

import (
	"detour/internal/application/shortener"
	"detour/internal/domain/url"
	"detour/internal/infrastructure/config"
	"detour/internal/infrastructure/http"
	"detour/internal/infrastructure/http/handlers"
	"detour/internal/infrastructure/mysql"
	"detour/internal/infrastructure/mysql/repositories"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize repository
	mysqlClient, err := mysql.NewClient(mysql.Config{
		Host:     cfg.MySQL.Host,
		Port:     cfg.MySQL.Port,
		User:     cfg.MySQL.User,
		Password: cfg.MySQL.Password,
		Database: cfg.MySQL.Database,
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
		Port:         cfg.Server.Port,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}, router)

	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
