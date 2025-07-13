package usecase

import (
	"article-test/internal/article/domain"
	"article-test/internal/article/dto"
	"article-test/internal/article/infrastructure"
	"article-test/internal/article/repository"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type ArticleService struct {
	repo  repository.ArticleRepository
	cache infrastructure.RedisArticleCache
}

func NewArticleService(r repository.ArticleRepository, c *redis.Client) *ArticleService {
	return &ArticleService{
		repo:  r,
		cache: *infrastructure.NewRedisArticleCache(c),
	}
}

func (s *ArticleService) Create(article *domain.Article) error {
	err := s.repo.Create(article)
	if err == nil {
		s.cache.ClearAllCachedArticles()
	}
	return err
}

func (s *ArticleService) GetAll(query, author string, limit, offset int) ([]dto.ArticleWithAuthorDTO, int, error) {
	cacheKey := fmt.Sprintf("articles:q=%s:a=%s:l=%d:o=%d", query, author, limit, offset)
	if data, ok := s.cache.GetCachedArticles(cacheKey); ok {
		return data, len(data), nil
	}

	articles, total, err := s.repo.GetAll(query, author, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	if len(articles) > 0 {
		s.cache.SetCachedArticles(cacheKey, articles)
	}

	return articles, total, nil
}
