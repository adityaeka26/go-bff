package usecase

import (
	"context"

	"github.com/adityaeka26/go-bff/services/order/internal/dto"
)

type OrderUsecase interface {
	GetOrders(ctx context.Context, request dto.GetOrdersRequest) ([]dto.GetOrdersResponse, error)
}
