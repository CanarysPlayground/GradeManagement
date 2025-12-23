package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/yourorg/grades-service/internal/models"
	"github.com/yourorg/grades-service/internal/repository"
	"github.com/yourorg/grades-service/internal/cache"
)

type GradeHandlers struct {
	repo repository.Repository
	cache cache.Cache
}

func NewGradeHandlers(r repository.Repository, c cache.Cache) *GradeHandlers {
	return &GradeHandlers{repo: r, cache: c}
}

func (h *GradeHandlers) ListGrades(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// try cache first
	if body, err := h.cache.Get(ctx, "grades:list"); err == nil && body != "" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
		return
	}
	grades, err := h.repo.ListGrades(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(grades)
	// set cache for 30s
	_ = h.cache.Set(ctx, "grades:list", string(b), 30*time.Second)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (h *GradeHandlers) CreateGrade(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var g models.Grade
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.repo.CreateGrade(ctx, g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// invalidate list cache
	_ = h.cache.Set(ctx, "grades:list", "", 1*time.Second)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *GradeHandlers) GetGrade(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	// try cache
	key := "grade:" + idStr
	if body, err := h.cache.Get(ctx, key); err == nil && body != "" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
		return
	}
	g, err := h.repo.GetGrade(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(g)
	_ = h.cache.Set(ctx, key, string(b), 60*time.Second)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (h *GradeHandlers) UpdateGrade(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var g models.Grade
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.repo.UpdateGrade(ctx, id, g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// invalidate caches
	_ = h.cache.Set(ctx, "grades:list", "", 1*time.Second)
	_ = h.cache.Set(ctx, "grade:"+idStr, "", 1*time.Second)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *GradeHandlers) DeleteGrade(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.repo.DeleteGrade(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// invalidate
	_ = h.cache.Set(ctx, "grades:list", "", 1*time.Second)
	_ = h.cache.Set(ctx, "grade:"+idStr, "", 1*time.Second)
	w.WriteHeader(http.StatusNoContent)
}
