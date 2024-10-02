package usecase

import (
	"context"

	"github.com/adityaeka26/go-bff/bff/web/internal/dto"
)

type WebUsecase interface {
	GetOrderHistory(ctx context.Context, request dto.GetOrderHistoryRequest) (*dto.GetOrderHistoryResponse, error)
}
