package cmd

import (
	"github.com/adityaeka26/go-bff/services/order/cmd/grpc"
	"github.com/adityaeka26/go-bff/services/order/config"
	"github.com/adityaeka26/go-bff/services/order/internal/model"
	"github.com/adityaeka26/go-bff/services/order/internal/usecase"
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

	err = postgres.GetDb().AutoMigrate(&model.Order{})
	if err != nil {
		logger.GetLog().Error("postgres migrate fail", zap.Error(err))
		panic(err)
	}

	if config.AppEnv != "production" {
		postgres.GetDb().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"user_id", "item", "quantity"}),
		}).Create(&model.Order{
			ID:       1,
			UserId:   1,
			Item:     "RTX 3060",
			Quantity: 2,
		})

		postgres.GetDb().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"user_id", "item", "quantity"}),
		}).Create(&model.Order{
			ID:       2,
			UserId:   1,
			Item:     "RTX 4060",
			Quantity: 1,
		})

		postgres.GetDb().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"user_id", "item", "quantity"}),
		}).Create(&model.Order{
			ID:       3,
			UserId:   2,
			Item:     "Avoskin Toner",
			Quantity: 2,
		})
	}

	orderUsecase := usecase.NewOrderUsecase(logger, postgres, config)

	err = grpc.ServeGRPC(config, orderUsecase)
	if err != nil {
		logger.GetLog().Error("serve grpc fail", zap.Error(err))
		panic(err)
	}
}
