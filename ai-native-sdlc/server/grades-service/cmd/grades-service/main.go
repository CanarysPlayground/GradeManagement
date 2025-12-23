package main

import (
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yourorg/grades-service/internal/handlers"
	"github.com/yourorg/grades-service/internal/repository"
	"github.com/yourorg/grades-service/internal/cache"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	redisAddr := os.Getenv("REDIS_ADDR")

	repo, err := repository.NewPostgresRepository(dsn)
	if err != nil { log.Fatal(err) }
	c := cache.NewRedisCache(redisAddr)

	h := handlers.NewGradeHandlers(repo, c)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/grades", h.ListGrades)
	r.Post("/grades", h.CreateGrade)
	r.Get("/grades/{id}", h.GetGrade)
	r.Put("/grades/{id}", h.UpdateGrade)
	r.Delete("/grades/{id}", h.DeleteGrade)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
