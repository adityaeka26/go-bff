package repository

import (
	"context"

	pbOrder "github.com/adityaeka26/go-bff/services/bff/internal/repository/order_proto"
	pbUser "github.com/adityaeka26/go-bff/services/bff/internal/repository/user_proto"
)

type UserRepository interface {
	GetUserInfo(ctx context.Context, id int) (*pbUser.GetUserInfoResponse, error)
}

type OrderRepository interface {
	GetOrders(ctx context.Context, userId int) (*pbOrder.GetOrdersResponse, error)
}
