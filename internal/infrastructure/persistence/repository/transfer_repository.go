package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormTransferRepository struct {
	db *gorm.DB
}

func NewGormTransferRepository(db *gorm.DB) repository.TransferRepository {
	return &GormTransferRepository{
		db: db,
	}
}

// CreateTransfer implements repository.TransferRepository.
func (g *GormTransferRepository) CreateTransfer(ctx context.Context, transfer *entity.Transfer) (*entity.Transfer, error) {
	model := &po.Transfer{
		FromAccountId: transfer.FromAccountId,
		ToAccountId:   transfer.ToAccountId,
		Amount:        transfer.Amount,
	}

	if err := g.db.WithContext(ctx).Model(&po.Transfer{}).Clauses(clause.Returning{}).Create(model).Error; err != nil {
		return nil, err
	}

	return &entity.Transfer{
		FromAccountId: model.FromAccountId,
		ToAccountId:   model.ToAccountId,
		Amount:        model.Amount,
	}, nil
}
