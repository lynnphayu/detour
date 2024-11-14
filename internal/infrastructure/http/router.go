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
	mux.HandleFunc("POST /api/v1/urls", urlHandler.ShortenURL)
	mux.HandleFunc("GET /{short}", urlHandler.RedirectURL)
	mux.HandleFunc("GET /api/v1/urls/{short}", urlHandler.GetURLDetails)

	// Apply middleware chain
	return chain.Then(mux)
}
