package service

import (
	"github.com/leontofel/link-shortener/internal/repository"
	"github.com/leontofel/link-shortener/pkg"
)

type ShortenerService struct {
	repo repository.Repository
}

func NewShortenerService(r repository.Repository) *ShortenerService {
	return &ShortenerService{repo: r}
}

func (s *ShortenerService) Shorten(url string) (string, error) {
	code := pkg.GenerateShortCode()
	return code, s.repo.Save(code, url)
}

func (s *ShortenerService) Resolve(code string) (string, error) {
	return s.repo.Find(code)
}
