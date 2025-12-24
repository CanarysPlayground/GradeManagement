package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/yourorg/grades-service/internal/models"
)

// MockRepository is an in-memory implementation for development without a database
type MockRepository struct {
	mu     sync.RWMutex
	grades map[int]models.Grade
	nextID int
}

// NewMockRepository creates a new mock repository with some sample data
func NewMockRepository() Repository {
	m := &MockRepository{
		grades: make(map[int]models.Grade),
		nextID: 1,
	}
	// Add some sample data
	m.grades[1] = models.Grade{ID: 1, StudentID: 101, Course: "Math", Score: 95}
	m.grades[2] = models.Grade{ID: 2, StudentID: 102, Course: "Science", Score: 88}
	m.grades[3] = models.Grade{ID: 3, StudentID: 103, Course: "English", Score: 92}
	m.nextID = 4
	return m
}

func (m *MockRepository) ListGrades(ctx context.Context) ([]models.Grade, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var out []models.Grade
	for _, g := range m.grades {
		out = append(out, g)
	}
	return out, nil
}

func (m *MockRepository) GetGrade(ctx context.Context, id int) (models.Grade, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	g, ok := m.grades[id]
	if !ok {
		return models.Grade{}, errors.New("grade not found")
	}
	return g, nil
}

func (m *MockRepository) CreateGrade(ctx context.Context, g models.Grade) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	g.ID = m.nextID
	m.grades[m.nextID] = g
	m.nextID++
	return g.ID, nil
}

func (m *MockRepository) UpdateGrade(ctx context.Context, id int, g models.Grade) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.grades[id]; !ok {
		return errors.New("grade not found")
	}
	g.ID = id
	m.grades[id] = g
	return nil
}

func (m *MockRepository) DeleteGrade(ctx context.Context, id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.grades[id]; !ok {
		return errors.New("grade not found")
	}
	delete(m.grades, id)
	return nil
}
