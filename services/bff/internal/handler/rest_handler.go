package handler

import (
	"github.com/adityaeka26/go-bff/services/bff/config"
	"github.com/adityaeka26/go-bff/services/bff/internal/dto"
	"github.com/adityaeka26/go-bff/services/bff/internal/usecase"
	pkgError "github.com/adityaeka26/go-pkg/error"
	"github.com/adityaeka26/go-pkg/helper"
	pkgValidator "github.com/adityaeka26/go-pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type restHandler struct {
	validator  *pkgValidator.XValidator
	webUsecase usecase.WebUsecase
}

func InitRestHandler(app *fiber.App, config *config.EnvConfig, validator *pkgValidator.XValidator, webUsecase usecase.WebUsecase) {
	handler := &restHandler{
		validator:  validator,
		webUsecase: webUsecase,
	}

	app.Get("/web/order/user/:user_id", handler.GetOrderHistory)
}

func (h *restHandler) GetOrderHistory(c *fiber.Ctx) error {
	req := &dto.GetOrderHistoryRequest{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	resp, err := h.webUsecase.GetOrderHistory(c.Context(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, resp, "get order history success")
}
