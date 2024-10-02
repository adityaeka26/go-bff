package cmd

import (
	"github.com/adityaeka26/go-bff/services/bff/cmd/rest"
	"github.com/adityaeka26/go-bff/services/bff/config"
	"github.com/adityaeka26/go-bff/services/bff/internal/repository"
	"github.com/adityaeka26/go-bff/services/bff/internal/usecase"
	"github.com/adityaeka26/go-pkg/logger"
	"go.uber.org/zap"
)

func Execute() {
	logger := logger.NewLogger()

	config, err := config.Load(".env")
	if err != nil {
		logger.GetLog().Error("load config fail", zap.Error(err))
		panic(err)
	}

	orderRepository, err := repository.NewOrderRepository(config)
	if err != nil {
		logger.GetLog().Error("init order repository fail", zap.Error(err))
		panic(err)
	}

	userRepository, err := repository.NewUserRepository(config)
	if err != nil {
		logger.GetLog().Error("init user repository fail", zap.Error(err))
		panic(err)
	}

	webUsecase := usecase.NewWebUsecase(logger, config, userRepository, orderRepository)

	err = rest.ServeREST(logger, config, webUsecase)
	if err != nil {
		logger.GetLog().Error("serve rest fail", zap.Error(err))
		panic(err)
	}
}
