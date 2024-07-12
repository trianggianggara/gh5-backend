package vote

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
	g.GET("", h.Get)
	g.POST("/upvote", h.Upvote)
	g.PUT("/downvote", h.Downvote)
}

func (h *delivery) Get(c echo.Context) error {
	result, err := h.Factory.Usecase.User.Find(c.Request().Context())
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result, "Get all users success").Send(c)
}

func (h *delivery) Upvote(c echo.Context) error {
	payload := new(dto.UpvoteRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.Vote.Upvote(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *delivery) Downvote(c echo.Context) error {
	payload := new(dto.DownvoteRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.Vote.Downvote(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
