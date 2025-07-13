package repository

import (
	"article-test/internal/article/domain"
	"article-test/internal/article/dto"
	"article-test/internal/article/infrastructure"
	"database/sql"
)

type ArticleRepository interface {
	Create(article *domain.Article) error
	GetAll(query, author string, limit, offset int) ([]dto.ArticleWithAuthorDTO, int, error)
}

type Article struct {
	infra infrastructure.PgArticleRepository
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &Article{
		infra: *infrastructure.NewPgArticleRepository(db),
	}
}

func (a *Article) Create(article *domain.Article) error {
	return a.infra.Create(article)
}

func (a *Article) GetAll(query string, author string, limit int, offset int) ([]dto.ArticleWithAuthorDTO, int, error) {
	return a.infra.GetAll(query, author, limit, offset)
}
