package cmd

import (
	"github.com/adityaeka26/go-bff/bff/web/cmd/rest"
	"github.com/adityaeka26/go-bff/bff/web/config"
	"github.com/adityaeka26/go-bff/bff/web/internal/repository"
	"github.com/adityaeka26/go-bff/bff/web/internal/usecase"
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

	userRepository, err := repository.NewUserRepository(config)
	if err != nil {
		logger.GetLog().Error("init user repository fail", zap.Error(err))
		panic(err)
	}

	webUsecase := usecase.NewWebUsecase(logger, config, userRepository)

	err = rest.ServeREST(logger, config, webUsecase)
	if err != nil {
		logger.GetLog().Error("serve rest fail", zap.Error(err))
		panic(err)
	}
}
