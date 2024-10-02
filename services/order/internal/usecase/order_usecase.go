package usecase

import (
	"context"

	"github.com/adityaeka26/go-bff/services/order/config"
	"github.com/adityaeka26/go-bff/services/order/internal/dto"
	"github.com/adityaeka26/go-bff/services/order/internal/model"
	"github.com/adityaeka26/go-pkg/logger"
	"github.com/adityaeka26/go-pkg/postgres"
	"go.uber.org/zap"
)

type orderUsecase struct {
	logger   *logger.Logger
	postgres *postgres.Postgres
	config   *config.EnvConfig
}

func NewOrderUsecase(logger *logger.Logger, postgres *postgres.Postgres, config *config.EnvConfig) OrderUsecase {
	return &orderUsecase{
		logger:   logger,
		postgres: postgres,
		config:   config,
	}
}

func (u *orderUsecase) GetOrders(ctx context.Context, request dto.GetOrdersRequest) ([]dto.GetOrdersResponse, error) {
	logger := u.logger.GetLog().With(zap.String("operationName", "orderUsecase.GetOrders"))

	var order []model.Order
	tx := u.postgres.GetDb().Where("user_id = ?", request.UserId).Find(&order)
	if tx.Error != nil {
		return nil, tx.Error
	}

	resp := []dto.GetOrdersResponse{}
	for _, o := range order {
		resp = append(resp, dto.GetOrdersResponse{
			OrderId:  int(o.ID),
			UserId:   int(o.UserId),
			Item:     o.Item,
			Quantity: int(o.Quantity),
		})
	}

	logger.Info("get orders success")
	return resp, nil
}
