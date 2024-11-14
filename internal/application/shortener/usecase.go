package shortener

import (
	"context"
	"detour/internal/domain/hit"
	"detour/internal/domain/url"
	"time"
)

type UseCase struct {
	urlService *url.Service
	hitService *hit.Service
}

func NewUseCase(urlService *url.Service, hitService *hit.Service) *UseCase {
	return &UseCase{
		urlService: urlService,
		hitService: hitService,
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

func (uc *UseCase) GetByShortURL(ctx context.Context, shortURL string) (*url.URL, error) {
	url, err := uc.urlService.GetByShortURL(ctx, shortURL)
	if err != nil {
		return nil, err
	}
	hits, err := uc.hitService.GetByURLID(ctx, url.ID)
	if err != nil {
		return nil, err
	}
	url.Hits = hits
	return url, nil
}

func (uc *UseCase) GetUrlToRedirect(ctx context.Context, shortURL string, header *HeaderDTO) (*URLResponseDTO, error) {
	url, err := uc.urlService.GetByShortURL(ctx, shortURL)
	if err != nil {
		return nil, err
	}
	uc.hitService.SaveHit(ctx, &hit.Hit{
		URLID:     url.ID,
		IP:        header.IP,
		UserAgent: header.UserAgent,
		Referer:   header.Referer,
		HitAt:     time.Now(),
	})

	return &URLResponseDTO{
		ShortURL:    url.Short,
		OriginalURL: url.Original,
	}, nil
}
