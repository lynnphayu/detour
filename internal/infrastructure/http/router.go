package http

import (
	"detour/internal/infrastructure/http/handlers"
	"detour/internal/infrastructure/http/middleware"
	"net/http"
)

func Setup(urlHandler *handlers.URLHandler) http.Handler {
	// Create middleware chain
	chain := middleware.NewChain(
		middleware.Recovery,
		middleware.Logging,
	)

	// Create router
	mux := http.NewServeMux()

	// Add routes
	mux.HandleFunc("/api/v1/urls", urlHandler.ShortenURL)
	mux.HandleFunc("/{short}", urlHandler.RedirectURL)

	// Apply middleware chain
	return chain.Then(mux)
}
