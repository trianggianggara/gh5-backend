package api

import (
	"gh5-backend/internal/delivery/api/auth"
	"gh5-backend/internal/delivery/api/cases"
	"gh5-backend/internal/delivery/api/lawyer"
	"gh5-backend/internal/delivery/api/role"
	"gh5-backend/internal/delivery/api/user"
	"gh5-backend/internal/delivery/api/vote"

	"gh5-backend/internal/factory"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	prefix := "api"

	auth.NewDelivery(f).Route(e.Group(prefix + "/auth"))
	cases.NewDelivery(f).Route(e.Group(prefix + "/cases"))
	user.NewDelivery(f).Route(e.Group(prefix + "/users"))
	lawyer.NewDelivery(f).Route(e.Group(prefix + "/lawyers"))
	role.NewDelivery(f).Route(e.Group(prefix + "/roles"))
	vote.NewDelivery(f).Route(e.Group(prefix + "/votes"))
}
