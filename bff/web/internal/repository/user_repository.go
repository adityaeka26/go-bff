package repository

import (
	"context"

	"github.com/adityaeka26/go-bff/bff/web/config"
	pbUser "github.com/adityaeka26/go-bff/bff/web/internal/repository/user_proto"
	pkgError "github.com/adityaeka26/go-pkg/error"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userRepository struct {
	config            *config.EnvConfig
	userServiceClient pbUser.UserServiceClient
}

func NewUserRepository(config *config.EnvConfig) (UserRepository, error) {
	conn, err := grpc.NewClient(config.UserServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pbUser.NewUserServiceClient(conn)

	return &userRepository{
		config:            config,
		userServiceClient: userServiceClient,
	}, nil
}

func (r *userRepository) GetUserInfo(ctx context.Context, id int) (*pbUser.GetUserInfoResponse, error) {
	res, err := r.userServiceClient.GetUserInfo(ctx, &pbUser.GetUserInfoRequest{
		UserId: int32(id),
	})
	if err != nil {
		return nil, pkgError.HttpError(err)
	}

	return res, nil
}
