package application

import "github.com/flopeztancredi/markdown-collab-api/internal/health/domain"

type Service struct {
	serviceName    string
	serviceVersion string
}

func NewService(name, version string) *Service {
	return &Service{
		serviceName:    name,
		serviceVersion: version,
	}
}

func (s *Service) Check() *domain.HealthStatus {
	return domain.NewHealthStatus(s.serviceName, s.serviceVersion)
}
