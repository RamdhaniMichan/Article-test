-- +goose Up
-- +goose StatementBegin
INSERT INTO
    authors (id, name)
VALUES (
        '11111111-1111-1111-1111-111111111111',
        'Alice Johnson'
    ),
    (
        '22222222-2222-2222-2222-222222222222',
        'Bob Smith'
    ),
    (
        '33333333-3333-3333-3333-333333333333',
        'Charlie Brown'
    );

INSERT INTO
    articles (
        id,
        author_id,
        title,
        body,
        created_at
    )
VALUES (
        'aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1',
        '11111111-1111-1111-1111-111111111111',
        'Getting Started with Go',
        'This article introduces the basics of Go programming language.',
        CURRENT_TIMESTAMP - INTERVAL '2 days'
    ),
    (
        'aaaaaaa2-aaaa-aaaa-aaaa-aaaaaaaaaaa2',
        '22222222-2222-2222-2222-222222222222',
        'Building REST APIs with Go',
        'Learn how to build RESTful APIs in Go without using any framework.',
        CURRENT_TIMESTAMP - INTERVAL '1 day'
    ),
    (
        'aaaaaaa3-aaaa-aaaa-aaaa-aaaaaaaaaaa3',
        '33333333-3333-3333-3333-333333333333',
        'Go Concurrency Patterns',
        'An advanced look into concurrency using goroutines and channels.',
        CURRENT_TIMESTAMP
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd