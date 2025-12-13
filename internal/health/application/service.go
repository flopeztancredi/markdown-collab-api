package application

import "github.com/flopeztancredi/markdown-collab-api/internal/health/domain"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Check() *domain.HealthStatus {
	return domain.NewHealthStatus()
}
