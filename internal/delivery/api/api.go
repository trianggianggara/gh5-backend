package api

import (
	"gh5-backend/internal/delivery/api/auth"
	"gh5-backend/internal/delivery/api/user"
	"gh5-backend/internal/factory"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	prefix := "api"

	user.NewDelivery(f).Route(e.Group(prefix + "/users"))
	auth.NewDelivery(f).Route(e.Group(prefix + "/auth"))

}
