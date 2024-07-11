package main

import (
	"gh5-backend/internal/driver/db"
	"gh5-backend/internal/driver/http"
	"gh5-backend/internal/factory"
	constant "gh5-backend/pkg/constants"
	"gh5-backend/pkg/utils/env"

	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv(constant.ENV)
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("running on env " + ENV)
}

func main() {
	db.Init()
	f := factory.Init()
	http.Init(f)
}
