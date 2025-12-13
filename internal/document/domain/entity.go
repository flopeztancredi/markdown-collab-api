package domain

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID        uuid.UUID
	Title     string
	Content   []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewDocument(title string) *Document {
	now := time.Now().UTC()
	return &Document{
		ID:        uuid.New(),
		Title:     title,
		Content:   nil,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
