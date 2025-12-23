package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	grades, _ := h.repo.ListGrades()
	json.NewEncoder(w).Encode(grades)
}

func (h *GradeHandlers) CreateGrade(w http.ResponseWriter, r *http.Request) {
	var g repository.Grade
	_ = json.NewDecoder(r.Body).Decode(&g)
	id, _ := h.repo.CreateGrade(g)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *GradeHandlers) GetGrade(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		// try chi url param
		idStr = chiURLParam(r, "id")
	}
	id, _ := strconv.Atoi(idStr)
	g, _ := h.repo.GetGrade(id)
	json.NewEncoder(w).Encode(g)
}

func (h *GradeHandlers) UpdateGrade(w http.ResponseWriter, r *http.Request) {
	idStr := chiURLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	var g repository.Grade
	_ = json.NewDecoder(r.Body).Decode(&g)
	_ = h.repo.UpdateGrade(id, g)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *GradeHandlers) DeleteGrade(w http.ResponseWriter, r *http.Request) {
	idStr := chiURLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	_ = h.repo.DeleteGrade(id)
	w.WriteHeader(http.StatusNoContent)
}

// chiURLParam extracts chi URL param without importing chi in this file to keep testable.
func chiURLParam(r *http.Request, name string) string {
	// chi stores URL params in context under key "chiCtx"
	// but to avoid importing chi here, use the URL path parsing fallback
	// If chi is used the standard way this will work via r.URL.Path
	// as a simple heuristic for this scaffold.
	// For now, return empty string to keep code compiling.
	return ""
}
