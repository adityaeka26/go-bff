package usecase

import (
	"context"

	"github.com/adityaeka26/go-bff/services/user/config"
	"github.com/adityaeka26/go-bff/services/user/internal/dto"
	"github.com/adityaeka26/go-bff/services/user/internal/model"
	pkgError "github.com/adityaeka26/go-pkg/error"
	"github.com/adityaeka26/go-pkg/logger"
	"github.com/adityaeka26/go-pkg/postgres"
	"go.uber.org/zap"
)

type userUsecase struct {
	logger   *logger.Logger
	postgres *postgres.Postgres
	config   *config.EnvConfig
}

func NewUserUsecase(logger *logger.Logger, postgres *postgres.Postgres, config *config.EnvConfig) UserUsecase {
	return &userUsecase{
		logger:   logger,
		postgres: postgres,
		config:   config,
	}
}

func (u *userUsecase) GetUserInfo(ctx context.Context, request dto.GetUserInfoRequest) (*dto.GetUserInfoResponse, error) {
	logger := u.logger.GetLog().With(zap.String("operationName", "userUsecase.GetUserInfo"))

	var user model.User
	tx := u.postgres.GetDb().First(&user, "id = ?", request.UserId)
	if tx.Error != nil {
		if tx.Error.Error() == "record not found" {
			logger.Warn("book not found")
			return nil, pkgError.BadRequest("book not found")
		}
		return nil, tx.Error
	}

	logger.Info("get user info success")
	return &dto.GetUserInfoResponse{
		UserId:   int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Location: user.Location,
	}, nil
}
