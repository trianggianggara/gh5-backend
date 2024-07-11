package repository

import (
	"gh5-backend/internal/model/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LawyerRepository struct {
	Repository[entity.LawyerModel]
	Log *logrus.Logger
}

func NewLawyerRepository(conn *gorm.DB, log *logrus.Logger) *LawyerRepository {
	model := entity.LawyerModel{}
	repository := NewRepository(conn, model, model.TableName())
	return &LawyerRepository{
		Repository: *repository,
		Log:        log,
	}
}
