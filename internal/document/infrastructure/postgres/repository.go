package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/flopeztancredi/markdown-collab-api/internal/document/domain"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, doc *domain.Document) error {
	query := `
		INSERT INTO documents (id, title, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(ctx, query, doc.ID, doc.Title, doc.CreatedAt, doc.UpdatedAt)
	return err
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Document, error) {
	query := `
		SELECT id, title, content, created_at, updated_at
		FROM documents
		WHERE id = $1
	`
	row := r.db.QueryRow(ctx, query, id)

	var doc domain.Document
	err := row.Scan(&doc.ID, &doc.Title, &doc.Content, &doc.CreatedAt, &doc.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *Repository) UpdateContent(ctx context.Context, id uuid.UUID, content []byte) error {
	query := `
		UPDATE documents
		SET content = $2, updated_at = $3
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, id, content, time.Now().UTC())
	return err
}
