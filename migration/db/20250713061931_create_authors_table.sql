-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE authors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name TEXT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE authors;
-- +goose StatementEnd