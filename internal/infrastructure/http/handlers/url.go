package handlers

import (
	"detour/internal/application/shortener"
	"detour/internal/infrastructure/http/response"
	"encoding/json"
	"fmt"
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

func (h *URLHandler) GetURLDetails(w http.ResponseWriter, r *http.Request) {
	shortURL := r.PathValue("short")
	fmt.Println(shortURL)
	if shortURL == "" {
		response.Error(w, http.StatusBadRequest, "INVALID_REQUEST", "Short URL is required")
		return
	}
	url, err := h.urlUseCase.GetByShortURL(r.Context(), shortURL)
	fmt.Println(err)
	if err != nil {
		response.Error(w, http.StatusNotFound, "URL_NOT_FOUND", "Short URL not found")
		return
	}

	response.JSON(w, http.StatusOK, url)
}

func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {

	header := &shortener.HeaderDTO{
		IP:        r.RemoteAddr,
		UserAgent: r.UserAgent(),
		Referer:   r.Header.Get("Referer"),
	}
	shortURL := r.URL.Path[1:]
	url, err := h.urlUseCase.GetUrlToRedirect(r.Context(), shortURL, header)

	if err != nil {
		response.Error(w, http.StatusNotFound, "URL_NOT_FOUND", "Short URL not found")
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusTemporaryRedirect)
}
