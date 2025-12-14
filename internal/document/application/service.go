package application

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/flopeztancredi/markdown-collab-api/internal/document/domain"
)

type Service struct {
	repo      domain.Repository
	dbTimeout time.Duration
}

func NewService(repo domain.Repository, dbTimeout time.Duration) *Service {
	return &Service{
		repo:      repo,
		dbTimeout: dbTimeout,
	}
}

func (s *Service) Create(ctx context.Context, title string) (*domain.Document, error) {
	ctx, cancel := context.WithTimeout(ctx, s.dbTimeout)
	defer cancel()

	doc := domain.NewDocument(title)
	if err := s.repo.Create(ctx, doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*domain.Document, error) {
	ctx, cancel := context.WithTimeout(ctx, s.dbTimeout)
	defer cancel()

	return s.repo.GetByID(ctx, id)
}

func (s *Service) UpdateContent(ctx context.Context, id uuid.UUID, content []byte) error {
	ctx, cancel := context.WithTimeout(ctx, s.dbTimeout)
	defer cancel()

	return s.repo.UpdateContent(ctx, id, content)
}
