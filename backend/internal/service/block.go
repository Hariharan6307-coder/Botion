package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
)

type BlockService struct {
	blockRepo *repository.BlockRepository
}

func NewBlockService(blockRepo *repository.BlockRepository) *BlockService {
	return &BlockService{blockRepo: blockRepo}
}

func (s *BlockService) CreateBlock(ctx context.Context, req models.CreateBlockRequest) (*models.Block, error) {
	return s.blockRepo.Create(ctx, req)
}

func (s *BlockService) GetBlock(ctx context.Context, id string) (*models.Block, error) {
    return s.blockRepo.GetByID(ctx, id)
}

func (s *BlockService) GetBlocksByPage(ctx context.Context, pageID string) ([]models.Block, error) {
    return s.blockRepo.GetByPageID(ctx, pageID)
}

func (s *BlockService) UpdateBlock(ctx context.Context, id string, req models.UpdateBlockRequest) (*models.Block, error) {
    return s.blockRepo.Update(ctx, id, req)
}

func (s *BlockService) DeleteBlock(ctx context.Context, id string) error {
    return s.blockRepo.Delete(ctx, id)
}
