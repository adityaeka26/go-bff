package rest

import (
	"fmt"
	"time"

	"github.com/adityaeka26/go-bff/services/bff/config"
	"github.com/adityaeka26/go-bff/services/bff/internal/handler"
	"github.com/adityaeka26/go-bff/services/bff/internal/usecase"
	"github.com/adityaeka26/go-pkg/graceful_shutdown"
	"github.com/adityaeka26/go-pkg/logger"
	pkgValidator "github.com/adityaeka26/go-pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ServeREST(logger *logger.Logger, config *config.EnvConfig, webUsecase usecase.WebUsecase) error {
	gs := graceful_shutdown.GracefulShutdown{
		Timeout:        5 * time.Second,
		GracefulPeriod: time.Duration(config.GracefulPeriod) * time.Second,
	}

	app := fiber.New()
	gs.Enable(app)
	gs.Register(logger)

	handler.InitRestHandler(app, config, &pkgValidator.XValidator{Validator: &validator.Validate{}}, webUsecase)

	app.Listen(fmt.Sprintf(":%s", config.RestPort))

	gs.Cleanup()

	return nil
}
