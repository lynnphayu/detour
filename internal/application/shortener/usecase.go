package shortener

import (
	"context"
	"detour/internal/domain/url"
)

type UseCase struct {
	urlService *url.Service
}

func NewUseCase(urlService *url.Service) *UseCase {
	return &UseCase{
		urlService: urlService,
	}
}

func (uc *UseCase) ShortenURL(ctx context.Context, dto *CreateURLDTO) (*URLResponseDTO, error) {
	url, err := uc.urlService.CreateShortURL(ctx, dto.OriginalURL)
	if err != nil {
		return nil, err
	}

	return &URLResponseDTO{
		ShortURL:    url.Short,
		OriginalURL: url.Original,
		CreatedAt:   url.CreatedAt,
	}, nil
}

func (uc *UseCase) GetUrlToRedirect(ctx context.Context, shortURL string, ip string) (*URLResponseDTO, error) {
	url, err := uc.urlService.GetByShortURL(ctx, shortURL)
	if err != nil {
		return nil, err
	}
	uc.urlService.IncrementHits(ctx, url.ID, ip)

	return &URLResponseDTO{
		ShortURL:    url.Short,
		OriginalURL: url.Original,
	}, nil
}
