package api

import (
	"gh5-backend/internal/delivery/api/auth"
	"gh5-backend/internal/delivery/api/cases"
	"gh5-backend/internal/delivery/api/role"
	"gh5-backend/internal/delivery/api/user"

	"gh5-backend/internal/factory"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	prefix := "api"

	auth.NewDelivery(f).Route(e.Group(prefix + "/auth"))
	cases.NewDelivery(f).Route(e.Group(prefix + "/cases"))
	user.NewDelivery(f).Route(e.Group(prefix + "/users"))
	role.NewDelivery(f).Route(e.Group(prefix + "/roles"))
}
