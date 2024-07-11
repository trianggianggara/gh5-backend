package middleware

import (
	"net/http"

	res "gh5-backend/pkg/utils/response"
	"gh5-backend/pkg/utils/validator"

	"github.com/labstack/echo/v4"

	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func ErrorHandler(err error, c echo.Context) {
	var errCustom *res.Error

	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	switch report.Code {
	case http.StatusNotFound:
		errCustom = res.ErrorBuilder(&res.ErrorConstant.RouteNotFound, err)
	default:
		errCustom = res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	res.ErrorResponse(errCustom).Send(c)
}

func Init(e *echo.Echo) {
	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
	)
	e.HTTPErrorHandler = ErrorHandler
	e.Validator = &validator.CustomValidator{Validator: validator.NewValidator()}
}
