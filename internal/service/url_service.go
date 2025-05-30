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

func NewURLService(r domain.URLRepository) *URLService {
	return &URLService{repo: r}
}

// Shorten generates a code, stores and returns the shortened URL
func (s *URLService) Shorten(original string) (*domain.URL, error) {
	// generate deterministic code
	h := sha1.Sum([]byte(original + fmt.Sprint(time.Now().UnixNano())))
	code := fmt.Sprintf("%x", h)[:8]
	u := &domain.URL{Code: code, Original: original}

	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *URLService) Resolve(code string) (*domain.URL, error) {
	return s.repo.FindByCode(code)
}
