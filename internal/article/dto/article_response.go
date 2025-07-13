package dto

import (
	"time"

	"github.com/google/uuid"
)

type ArticleWithAuthorDTO struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    AuthorDTO `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}
