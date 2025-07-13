-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    author_id UUID NOT NULL,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE CASCADE
);

CREATE INDEX idx_articles_fts ON articles USING GIN (
    to_tsvector (
        'english',
        title || ' ' || body
    )
);

CREATE INDEX idx_articles_author_id ON articles (author_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE articles;
-- +goose StatementEnd