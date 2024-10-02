package repository

import (
	"context"

	"github.com/adityaeka26/go-bff/services/bff/config"
	pbOrder "github.com/adityaeka26/go-bff/services/bff/internal/repository/order_proto"
	pkgError "github.com/adityaeka26/go-pkg/error"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type orderRepository struct {
	config             *config.EnvConfig
	orderServiceClient pbOrder.OrderServiceClient
}

func NewOrderRepository(config *config.EnvConfig) (OrderRepository, error) {
	conn, err := grpc.NewClient(config.OrderServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	orderServiceClient := pbOrder.NewOrderServiceClient(conn)

	return &orderRepository{
		config:             config,
		orderServiceClient: orderServiceClient,
	}, nil
}

func (r *orderRepository) GetOrders(ctx context.Context, userId int) (*pbOrder.GetOrdersResponse, error) {
	res, err := r.orderServiceClient.GetOrders(ctx, &pbOrder.GetOrdersRequest{
		UserId: int32(userId),
	})
	if err != nil {
		return nil, pkgError.HttpError(err)
	}

	return res, nil
}
