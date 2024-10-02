package usecase

import (
	"context"

	"github.com/adityaeka26/go-bff/services/bff/config"
	"github.com/adityaeka26/go-bff/services/bff/internal/dto"
	"github.com/adityaeka26/go-bff/services/bff/internal/repository"
	"github.com/adityaeka26/go-pkg/logger"
)

type webUsecase struct {
	logger          *logger.Logger
	config          *config.EnvConfig
	userRepository  repository.UserRepository
	orderRepository repository.OrderRepository
}

func NewWebUsecase(logger *logger.Logger, config *config.EnvConfig, userRepository repository.UserRepository, orderRepository repository.OrderRepository) WebUsecase {
	return &webUsecase{
		logger:          logger,
		config:          config,
		userRepository:  userRepository,
		orderRepository: orderRepository,
	}
}

func (u *webUsecase) GetOrderHistory(ctx context.Context, request dto.GetOrderHistoryRequest) (*dto.GetOrderHistoryResponse, error) {
	user, err := u.userRepository.GetUserInfo(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	orders, err := u.orderRepository.GetOrders(ctx, int(user.UserId))
	if err != nil {
		return nil, err
	}

	ordersResp := []dto.GetOrderHistoryResponseOrders{}
	for _, order := range orders.Orders {
		ordersResp = append(ordersResp, dto.GetOrderHistoryResponseOrders{
			OrderId:  int(order.OrderId),
			Item:     order.Item,
			Quantity: int(order.Quantity),
		})
	}

	return &dto.GetOrderHistoryResponse{
		UserId:   int(user.UserId),
		Name:     user.Name,
		Email:    user.Email,
		Location: user.Location,
		Orders:   ordersResp,
	}, nil
}
