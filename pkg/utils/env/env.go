package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var envFile string

type Env interface {
	GetString(name string) string
	GetBool(name string) bool
	GetInt(name string) int
	GetFloat(name string) float64
}

type env struct{}

func NewEnv() *env {
	return &env{}
}

func (e *env) Load(env string) {
	switch env {
	case "STG":
		envFile = ".env.staging"
	case "PROD":
		envFile = ".env.production"
	case "DEV":
		envFile = ".env.development"
	default:
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		logrus.Errorf("Load .env file error: %s, cause: %v", err, err)
		os.Exit(-1)
	}
}

func (e *env) GetString(name string) string {
	return os.Getenv(name)
}

func (e *env) GetBool(name string) bool {
	s := e.GetString(name)
	i, err := strconv.ParseBool(s)
	if nil != err {
		return false
	}
	return i
}

func (e *env) GetInt(name string) int {
	s := e.GetString(name)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func (e *env) GetFloat(name string) float64 {
	s := e.GetString(name)
	i, err := strconv.ParseFloat(s, 64)
	if nil != err {
		return 0
	}
	return i
}

func (e *env) GetFile() string {
	return envFile
}
