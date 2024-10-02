package usecase

import (
	"context"
	"fmt"

	"github.com/adityaeka26/go-bff/bff/web/config"
	"github.com/adityaeka26/go-bff/bff/web/internal/dto"
	"github.com/adityaeka26/go-bff/bff/web/internal/repository"
	"github.com/adityaeka26/go-pkg/logger"
)

type webUsecase struct {
	logger         *logger.Logger
	config         *config.EnvConfig
	userRepository repository.UserRepository
}

func NewWebUsecase(logger *logger.Logger, config *config.EnvConfig, userRepository repository.UserRepository) WebUsecase {
	return &webUsecase{
		logger:         logger,
		config:         config,
		userRepository: userRepository,
	}
}

func (u *webUsecase) GetOrderHistory(ctx context.Context, request dto.GetOrderHistoryRequest) (*dto.GetOrderHistoryResponse, error) {
	user, err := u.userRepository.GetUserInfo(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	return &dto.GetOrderHistoryResponse{
		UserId:   int(user.UserId),
		Name:     user.Name,
		Email:    user.Email,
		Location: user.Location,
		Orders:   []string{},
	}, nil
}
