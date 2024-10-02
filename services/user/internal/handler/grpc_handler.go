package handler

import (
	"context"

	"github.com/adityaeka26/go-bff/services/user/internal/dto"
	pb "github.com/adityaeka26/go-bff/services/user/internal/handler/proto"
	"github.com/adityaeka26/go-bff/services/user/internal/usecase"
	pkgError "github.com/adityaeka26/go-pkg/error"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedUserServiceServer
	userUsecase usecase.UserUsecase
}

func InitGrpcHandler(app *grpc.Server, userUsecase usecase.UserUsecase) {
	handler := &grpcHandler{
		userUsecase: userUsecase,
	}

	pb.RegisterUserServiceServer(app, handler)
}

func (h *grpcHandler) GetUserInfo(ctx context.Context, request *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	response, err := h.userUsecase.GetUserInfo(ctx, dto.GetUserInfoRequest{
		UserId: int(request.UserId),
	})
	if err != nil {
		return nil, pkgError.GrpcError(err)
	}

	return &pb.GetUserInfoResponse{
		UserId:   int32(response.UserId),
		Name:     response.Name,
		Email:    response.Email,
		Location: response.Location,
	}, nil
}
