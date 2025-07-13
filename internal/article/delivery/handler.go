package delivery

import (
	"article-test/internal/article/domain"
	"article-test/internal/article/usecase"
	"article-test/pkg/utils"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	service *usecase.ArticleService
}

func NewArticleHandler(s *usecase.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: s}
}

func (h *ArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		h.list(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ArticleHandler) create(w http.ResponseWriter, r *http.Request) {
	var article domain.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Error(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.service.Create(&article); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.Success(
		http.StatusCreated,
		"Article created successfully",
		article,
		nil,
	))
}

func (h *ArticleHandler) list(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	author := r.URL.Query().Get("author")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if limit == 0 {
		limit = 100
	}

	if offset == 0 {
		offset = 0
	}

	articles, total, err := h.service.GetAll(query, author, limit, offset)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	currentPage := (offset / limit) + 1
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	meta := utils.PaginationMeta{
		CurrentPage:  currentPage,
		TotalPages:   totalPages,
		TotalRecords: total,
		Limit:        limit,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.Success(
		http.StatusOK,
		"Articles retrieved successfully",
		articles,
		meta,
	))
}
