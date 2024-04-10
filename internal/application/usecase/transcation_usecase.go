package usecase

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/dto"
)

type TransactionUseCase interface {
	Transfer(context.Context, *dto.TransferRequest) (*dto.TransferResponse, error)
	TransferWithTrx(context.Context, *dto.TransferRequest) (*dto.TransferResponse, error)
}
