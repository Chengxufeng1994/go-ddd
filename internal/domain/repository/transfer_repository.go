package repository

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
)

type TransferRepository interface {
	CreateTransfer(context.Context, *entity.Transfer) (*entity.Transfer, error)
}
