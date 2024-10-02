package cmd

import (
	"github.com/adityaeka26/go-bff/services/user/cmd/grpc"
	"github.com/adityaeka26/go-bff/services/user/config"
	"github.com/adityaeka26/go-bff/services/user/internal/model"
	"github.com/adityaeka26/go-bff/services/user/internal/usecase"
	"github.com/adityaeka26/go-pkg/logger"
	"github.com/adityaeka26/go-pkg/postgres"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

func Execute() {
	logger := logger.NewLogger()

	config, err := config.Load(".env")
	if err != nil {
		logger.GetLog().Error("load config fail", zap.Error(err))
		panic(err)
	}

	postgres, err := postgres.NewPostgres(
		config.PostgresUsername,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresDb,
		config.PostgresSslEnabled,
	)
	if err != nil {
		logger.GetLog().Error("init postgres fail", zap.Error(err))
		panic(err)
	}

	err = postgres.GetDb().AutoMigrate(&model.User{})
	if err != nil {
		logger.GetLog().Error("postgres migrate fail", zap.Error(err))
		panic(err)
	}

	if config.AppEnv != "production" {
		postgres.GetDb().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "email", "location"}),
		}).Create(&model.User{
			ID:       1,
			Name:     "Aditya",
			Email:    "aditya@adityaeka.my.id",
			Location: "Jakarta, Indonesia",
		})

		postgres.GetDb().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "email", "location"}),
		}).Create(&model.User{
			ID:       2,
			Name:     "Eka",
			Email:    "eka@adityaeka.my.id",
			Location: "Tulungagung, Indonesia",
		})

		postgres.GetDb().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "email", "location"}),
		}).Create(&model.User{
			ID:       3,
			Name:     "Bagas",
			Email:    "bagas@adityaeka.my.id",
			Location: "Tulungagung, Indonesia",
		})
	}

	userUsecase := usecase.NewUserUsecase(logger, postgres, config)

	err = grpc.ServeGRPC(config, userUsecase)
	if err != nil {
		logger.GetLog().Error("serve grpc fail", zap.Error(err))
		panic(err)
	}
}
