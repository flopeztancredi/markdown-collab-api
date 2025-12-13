package domain

import "time"

type HealthStatus struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

func NewHealthStatus(service, version string) *HealthStatus {
	return &HealthStatus{
		Status:    "healthy",
		Service:   service,
		Version:   version,
		Timestamp: time.Now().UTC(),
	}
}
