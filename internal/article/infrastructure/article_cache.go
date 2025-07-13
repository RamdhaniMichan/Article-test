package infrastructure

import (
	"article-test/internal/article/dto"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisArticleCache struct {
	client *redis.Client
	ctx    context.Context
	ttl    time.Duration
}

func NewRedisArticleCache(rdb *redis.Client) *RedisArticleCache {
	return &RedisArticleCache{
		client: rdb,
		ctx:    context.Background(),
		ttl:    15 * time.Minute,
	}
}

func (r *RedisArticleCache) GetCachedArticles(key string) ([]dto.ArticleWithAuthorDTO, bool) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return nil, false
	}

	var articles []dto.ArticleWithAuthorDTO
	if err := json.Unmarshal([]byte(val), &articles); err != nil {
		return nil, false
	}
	return articles, true
}

func (r *RedisArticleCache) SetCachedArticles(key string, articles []dto.ArticleWithAuthorDTO) {
	data, _ := json.Marshal(articles)
	r.client.Set(r.ctx, key, data, r.ttl)
}

func (r *RedisArticleCache) ClearAllCachedArticles() {
	r.client.FlushDB(r.ctx)
}
