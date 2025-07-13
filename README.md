# Article Service API

A RESTful API for managing articles using Go and PostgreSQL.

---

## Technologies

- Go 1.23
- PostgreSQL 15 (Docker)
- Redis 7 (Docker)
- Docker & Docker Compose
- Goose (database migration)

---

## Getting Started

### 1. Clone the Repository

```sh
git clone https://github.com/yourusername/article-test.git
cd article-test
```

### 2. Configure Environment Variables

Copy `.env.example` to `.env` and adjust the values if needed:

```sh
cp .env.example .env
```

### 3. Run with Docker Compose

```sh
docker-compose up --build
```

- The API will be available at `http://localhost:8081`
- PostgreSQL will be available at `localhost:5433`
- Redis will be available at `localhost:63791`

### 4. Database Migration

Migrations will run automatically on startup. To run manually:

```sh
docker-compose exec app goose -dir ./migration/db postgres "host=$DB_HOST user=$DB_USER password=$DB_PASS dbname=$DB_NAME sslmode=disable" up
```

---

## API Endpoints

### GET /articles

Retrieve a list of articles (supports filtering and pagination):

- `query`: search in title/body
- `author`: filter by author name
- `limit`: items per page (default: 100)
- `offset`: offset for pagination

Example:
```
GET /articles?query=go&author=Alice&limit=10&offset=0
```

### POST /articles

Create a new article:

```json
{
  "title": "Article Title",
  "body": "Article content",
  "author_id": "uuid-author"
}
```

---

## Project Structure

```
.
├── cmd/                # Application entry point
├── config/             # Database and Redis configuration
├── internal/article/   # Article module (domain, usecase, repository, delivery, infrastructure)
├── migration/db/       # Database migration files
├── pkg/utils/          # Utilities (response, pagination)
├── Dockerfile
├── docker-compose.yaml
├── .env.example
└── README.md
```