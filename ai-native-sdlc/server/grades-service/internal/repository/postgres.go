package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/yourorg/grades-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository defines methods to interact with the database.
type Repository interface {
	ListGrades(ctx context.Context) ([]models.Grade, error)
	GetGrade(ctx context.Context, id int) (models.Grade, error)
	CreateGrade(ctx context.Context, g models.Grade) (int, error)
	UpdateGrade(ctx context.Context, id int, g models.Grade) error
	DeleteGrade(ctx context.Context, id int) error
}

type postgresRepo struct{
	pool *pgxpool.Pool
}

// NewPostgresRepository connects to Postgres using the provided DSN.
func NewPostgresRepository(dsn string) (Repository, error) {
	if dsn == "" {
		return nil, errors.New("DB_DSN not provided")
	}
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	cfg.MaxConns = 5
	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}
	return &postgresRepo{pool: pool}, nil
}

func (p *postgresRepo) ListGrades(ctx context.Context) ([]models.Grade, error) {
	rows, err := p.pool.Query(ctx, `SELECT id, student_id, course, score FROM grades ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []models.Grade
	for rows.Next() {
		var g models.Grade
		if err := rows.Scan(&g.ID, &g.StudentID, &g.Course, &g.Score); err != nil {
			return nil, err
		}
		out = append(out, g)
	}
	return out, nil
}

func (p *postgresRepo) GetGrade(ctx context.Context, id int) (models.Grade, error) {
	var g models.Grade
	err := p.pool.QueryRow(ctx, `SELECT id, student_id, course, score FROM grades WHERE id=$1`, id).Scan(&g.ID, &g.StudentID, &g.Course, &g.Score)
	if err != nil {
		return models.Grade{}, err
	}
	return g, nil
}

func (p *postgresRepo) CreateGrade(ctx context.Context, g models.Grade) (int, error) {
	var id int
	err := p.pool.QueryRow(ctx, `INSERT INTO grades (student_id, course, score) VALUES ($1,$2,$3) RETURNING id`, g.StudentID, g.Course, g.Score).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *postgresRepo) UpdateGrade(ctx context.Context, id int, g models.Grade) error {
	ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := p.pool.Exec(ctx2, `UPDATE grades SET student_id=$1, course=$2, score=$3 WHERE id=$4`, g.StudentID, g.Course, g.Score, id)
	return err
}

func (p *postgresRepo) DeleteGrade(ctx context.Context, id int) error {
	_, err := p.pool.Exec(context.Background(), `DELETE FROM grades WHERE id=$1`, id)
	return err
}

// helper used by cache to serialize grade
func gradeToJSON(g models.Grade) (string, error) {
	b, err := json.Marshal(g)
	return string(b), err
}
