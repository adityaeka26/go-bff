package usecase

import (
	"context"

	"github.com/adityaeka26/go-bff/services/user/internal/dto"
)

type UserUsecase interface {
	GetUserInfo(ctx context.Context, request dto.GetUserInfoRequest) (*dto.GetUserInfoResponse, error)
}
