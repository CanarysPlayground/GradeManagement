package repository

import "errors"

// Repository defines methods to interact with the database.
type Repository interface {
	ListGrades() ([]Grade, error)
	GetGrade(id int) (Grade, error)
	CreateGrade(g Grade) (int, error)
	UpdateGrade(id int, g Grade) error
	DeleteGrade(id int) error
}

// Grade model used by repository
type Grade struct {
	ID        int
	StudentID int
	Course    string
	Score     float64
}

// NewPostgresRepository returns a stub implementation for now.
func NewPostgresRepository(dsn string) (Repository, error) {
	if dsn == "" {
		return nil, errors.New("DB_DSN not provided")
	}
	return &postgresRepo{}, nil
}

type postgresRepo struct{}

func (p *postgresRepo) ListGrades() ([]Grade, error) { return []Grade{}, nil }
func (p *postgresRepo) GetGrade(id int) (Grade, error)   { return Grade{}, nil }
func (p *postgresRepo) CreateGrade(g Grade) (int, error) { return 0, nil }
func (p *postgresRepo) UpdateGrade(id int, g Grade) error { return nil }
func (p *postgresRepo) DeleteGrade(id int) error        { return nil }
