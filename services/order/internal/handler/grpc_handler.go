package handler

import (
	"context"

	"github.com/adityaeka26/go-bff/services/order/internal/dto"
	pb "github.com/adityaeka26/go-bff/services/order/internal/handler/proto"
	"github.com/adityaeka26/go-bff/services/order/internal/usecase"
	pkgError "github.com/adityaeka26/go-pkg/error"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	orderUsecase usecase.OrderUsecase
}

func InitGrpcHandler(app *grpc.Server, orderUsecase usecase.OrderUsecase) {
	handler := &grpcHandler{
		orderUsecase: orderUsecase,
	}

	pb.RegisterOrderServiceServer(app, handler)
}

func (h *grpcHandler) GetOrders(ctx context.Context, request *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	response, err := h.orderUsecase.GetOrders(ctx, dto.GetOrdersRequest{
		UserId: int(request.UserId),
	})
	if err != nil {
		return nil, pkgError.GrpcError(err)
	}

	orders := []*pb.Order{}
	for _, r := range response {
		orders = append(orders, &pb.Order{
			OrderId:  int32(r.OrderId),
			UserId:   int32(r.UserId),
			Item:     r.Item,
			Quantity: int32(r.Quantity),
		})
	}

	return &pb.GetOrdersResponse{
		Orders: orders,
	}, nil
}
