package http

import (
	"fmt"

	"gh5-backend/internal/delivery/api"
	"gh5-backend/internal/delivery/middleware"
	"gh5-backend/internal/factory"
	"gh5-backend/pkg/constants"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Init(f factory.Factory) {
	e := echo.New()

	middleware.Init(e)

	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", os.Getenv(constants.APP), os.Getenv(constants.VERSION))
		return c.String(http.StatusOK, message)
	})

	api.Init(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv(constants.PORT)))
}
