package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"url-shortner-go/internal/domain"
)

type URLService struct {
	repo domain.URLRepository
}

func NewURLService(repo domain.URLRepository) *URLService {
	return &URLService{repo: repo}
}

// Shorten generates a code, stores and returns the shortened URL
func (service *URLService) Shorten(original string) (*domain.URL, error) {
	// generate deterministic code
	generated := sha1.Sum([]byte(original + fmt.Sprint(time.Now().UnixNano())))
	code := fmt.Sprintf("%x", generated)[:8]
	url := &domain.URL{Code: code, Original: original}

	if err := service.repo.Create(url); err != nil {
		return nil, err
	}
	return url, nil
}

func (service *URLService) Resolve(code string) (*domain.URL, error) {
	return service.repo.FindByCode(code)
}
