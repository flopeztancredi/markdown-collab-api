package application

import (
	"context"

	"github.com/google/uuid"

	"github.com/flopeztancredi/markdown-collab-api/internal/document/domain"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, title string) (*domain.Document, error) {
	doc := domain.NewDocument(title)
	if err := s.repo.Create(ctx, doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*domain.Document, error) {
	return s.repo.GetByID(ctx, id)
}
