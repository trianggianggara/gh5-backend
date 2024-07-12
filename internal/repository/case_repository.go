package repository

import (
	model "gh5-backend/internal/model/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CaseRepository struct {
	Repository[model.CaseModel]
	Log *logrus.Logger
}

func NewCaseRepository(conn *gorm.DB, log *logrus.Logger) *CaseRepository {
	model := model.CaseModel{}
	repository := NewRepository(conn, model, model.TableName())
	return &CaseRepository{
		Repository: *repository,
		Log:        log,
	}
}
