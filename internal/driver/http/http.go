package http

import (
	"fmt"

	"gh5-backend/internal/delivery/api"
	"gh5-backend/internal/factory"
	"gh5-backend/pkg/constants"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/acme/autocert"
)

func Init(f factory.Factory) {
	e := echo.New()

	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", os.Getenv(constants.APP), os.Getenv(constants.VERSION))
		return c.String(http.StatusOK, message)
	})

	api.Init(e, f)

	e.Logger.Fatal(e.StartAutoTLS(":" + os.Getenv(constants.PORT)))
}
