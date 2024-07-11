package auth

import (
	"gh5-backend/internal/factory"
	"gh5-backend/internal/model/dto"
	res "gh5-backend/pkg/utils/response"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	Factory factory.Factory
}

func NewDelivery(f factory.Factory) *delivery {
	return &delivery{f}
}

func (h *delivery) Route(g *echo.Group) {
	g.POST("/login", h.Login)
	g.POST("/register", h.Register)
}

func (h *delivery) Login(c echo.Context) error {
	payload := new(dto.AuthLoginRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.Factory.Usecase.Auth.Login(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

func (h *delivery) Register(c echo.Context) error {
	payload := new(dto.AuthRegisterRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.Factory.Usecase.Auth.Register(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
