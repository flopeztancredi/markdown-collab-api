package domain

import "time"

type HealthStatus struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func NewHealthStatus() *HealthStatus {
	return &HealthStatus{
		Status:    "healthy",
		Timestamp: time.Now().UTC(),
	}
}
