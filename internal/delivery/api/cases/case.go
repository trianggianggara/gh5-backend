package cases

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
	g.GET("/:id", h.GetByID)
	g.POST("/", h.Create)
	g.PUT("/:id", h.UpdateByID)

}

func (h *delivery) Get(c echo.Context) error {
	result, err := h.Factory.Usecase.Case.Find(c.Request().Context())
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result, "Get all cases success").Send(c)
}

func (h *delivery) GetByID(c echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.Case.FindByID(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result, "Get case by id success").Send(c)
}

func (h *delivery) Create(c echo.Context) error {
	payload := new(dto.CreateCaseRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.Case.Create(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *delivery) UpdateByID(c echo.Context) error {
	payload := new(dto.UpdateCaseRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.Factory.Usecase.Case.UpdateByID(c.Request().Context(), *payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
