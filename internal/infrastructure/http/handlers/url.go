package handlers

import (
	"detour/internal/application/shortener"
	"detour/internal/infrastructure/http/response"
	"encoding/json"
	"net/http"
)

type URLHandler struct {
	urlUseCase *shortener.UseCase
}

// ServeHTTP implements http.Handler.
func (h *URLHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func NewURLHandler(urlUseCase *shortener.UseCase) *URLHandler {
	return &URLHandler{
		urlUseCase: urlUseCase,
	}
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
		return
	}

	var req shortener.CreateURLDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body")
		return
	}

	result, err := h.urlUseCase.ShortenURL(r.Context(), &req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "SHORTENING_FAILED", err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, result)
}

func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Error(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
		return
	}

	shortURL := r.URL.Path[1:] // Remove leading slash
	url, err := h.urlUseCase.GetUrlToRedirect(r.Context(), shortURL, r.RemoteAddr)

	if err != nil {
		response.Error(w, http.StatusNotFound, "URL_NOT_FOUND", "Short URL not found")
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusTemporaryRedirect)
}
