package repository

import (
	"backend/internal/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BlockRepository struct {
	db *pgxpool.Pool
}

func NewBlockRepository(db *pgxpool.Pool) *BlockRepository {
	return &BlockRepository{db: db}
}

func (r *BlockRepository) GetPage(ctx context.Context) (*models.Page, error){
	page := &models.Page{}
	query := `
		SELECT id, title
		FROM "Page"
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, "550e8400-e29b-41d4-a716-446655440000").Scan(
		&page.ID, &page.Title,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("Page not found")
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to get page: %w", err)
	}

	return page, nil
}


func (r *BlockRepository) Create(ctx context.Context, req models.CreateBlockRequest) (*models.Block, error) {
	block := &models.Block{
		ID:           uuid.New().String(),
		PageID: 	  req.PageID,
		Type:        req.Type,
		Content:     req.Content,
		Position:    req.Position,
		ParentBlockID: req.ParentBlockID,
	}

	query := `
		INSERT INTO "Block" (id, "pageId", type, content, position, "parentBlockId")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, "pageId", type, content, position, "parentBlockId"
	`

	err := r.db.QueryRow(ctx, query, 
		block.ID, block.PageID, block.Type, block.Content, block.Position, block.ParentBlockID,
	).Scan(
		&block.ID, &block.PageID, &block.Type, &block.Content, &block.Position, &block.ParentBlockID,
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to create block: %w", err)
	}

	return block, nil
} 

func (r *BlockRepository) GetByID(ctx context.Context, id string) (*models.Block, error) {
	block := &models.Block{}
	query := `
		SELECT id, "pageId", type, content, position, "parentBlockId"
		FROM "Block"
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&block.ID, &block.PageID, &block.Type, &block.Content, &block.Position, &block.ParentBlockID,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("Block not found")
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to get block by ID: %w", err)
	}

	return block, nil
}

func (r *BlockRepository) GetByPageID(ctx context.Context, pageID string) ([]models.Block, error) {
	query := `
		SELECT id, "pageId", type, content, position, "parentBlockId"
		FROM "Block"
		WHERE "pageId" = $1
		ORDER BY position ASC
	`

	rows, err := r.db.Query(ctx, query, pageID)

	if err != nil {
		return nil, fmt.Errorf("Failed to get blocks by page ID: %w", err)
	}
	defer rows.Close()

	var blocks []models.Block
	for rows.Next() {
		var block models.Block
		err := rows.Scan(
			&block.ID,
			&block.PageID,
			&block.Type,
			&block.Content,
			&block.Position,
			&block.ParentBlockID,
		)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan block: %w", err)
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

func (r *BlockRepository) Update(ctx context.Context, id string, req models.UpdateBlockRequest) (*models.Block, error) {
	query := `
		UPDATE "Block"
		SET type = COALESCE($1, type),
			content = COALESCE($2, content),
			position = COALESCE($3, position),
			"parentBlockId" = COALESCE($4, "parentBlockId")
		WHERE id = $5
		RETURNING id, "pageId", type, content, position, "parentBlockId"	
	`
	block := &models.Block{}
	err := r.db.QueryRow(ctx, query,
		id, req.Type, req.Content, req.Position, req.ParentBlockID,
	).Scan(
		&block.ID, &block.PageID, &block.Type, &block.Content, &block.Position, &block.ParentBlockID,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("Block not found")
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to update block: %w", err)
	}

	return block, nil
}

func (r *BlockRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM "Block"
		WHERE id = $1
	`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete block: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("Block not found")
	}

	return nil
}