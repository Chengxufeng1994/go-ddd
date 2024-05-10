package usecase

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
)

type TransactionUseCase interface {
	Transfer(context.Context, *dto.TransferRequest) (*dto.TransferResponse, error)
	TransferWithTrx(context.Context, *dto.TransferRequest) (*dto.TransferResponse, error)
}
