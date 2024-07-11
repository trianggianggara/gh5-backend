package factory

import (
	"gh5-backend/internal/factory/repository"
	"gh5-backend/internal/factory/usecase"
)

type Factory struct {
	Repository repository.Factory
	Usecase    usecase.Factory
}

func Init() Factory {
	f := Factory{}

	f.Repository = repository.Init()
	f.Usecase = usecase.Init(f.Repository)

	return f
}
