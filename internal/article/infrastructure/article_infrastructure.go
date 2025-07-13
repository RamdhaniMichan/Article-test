package infrastructure

import (
	"article-test/internal/article/domain"
	"article-test/internal/article/dto"
	"database/sql"
	"fmt"
)

type PgArticleRepository struct {
	db *sql.DB
}

func NewPgArticleRepository(db *sql.DB) *PgArticleRepository {
	return &PgArticleRepository{db: db}
}

func (r *PgArticleRepository) Create(article *domain.Article) error {
	query := `INSERT INTO articles (title, body, author_id) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, article.Title, article.Body, article.AuthorID)
	return err
}

func (r *PgArticleRepository) GetAll(query, author string, limit, offset int) ([]dto.ArticleWithAuthorDTO, int, error) {
	base := `
		SELECT a.id, a.title, a.body, a.created_at, a.author_id, au.name
		FROM articles a
		JOIN authors au ON a.author_id = au.id
	`

	conditions := []string{}
	args := []interface{}{}

	if query != "" {
		conditions = append(conditions, `to_tsvector(title || ' ' || body) @@ websearch_to_tsquery($1)`)
		args = append(args, query)
	}

	if author != "" {
		conditions = append(conditions, fmt.Sprintf("au.name = $%d", len(args)+1))
		args = append(args, author)
	}

	if len(conditions) > 0 {
		base += " WHERE " + conditions[0]
		for i := 1; i < len(conditions); i++ {
			base += " AND " + conditions[i]
		}
	}

	base += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, limit, offset)

	rows, err := r.db.Query(base, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var articles []dto.ArticleWithAuthorDTO
	for rows.Next() {
		var a dto.ArticleWithAuthorDTO
		if err := rows.Scan(&a.ID, &a.Title, &a.Body, &a.CreatedAt, &a.Author.ID, &a.Author.Name); err != nil {
			return nil, 0, err
		}
		articles = append(articles, a)
	}

	return articles, len(articles), nil
}
