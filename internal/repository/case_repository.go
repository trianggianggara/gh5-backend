package repository

import (
	"context"
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

func (r *Repository[T]) FindAll(ctx context.Context) ([]model.CaseModel, error) {
	r.checkTrx(ctx)
	query := r.getConn().Model(model.CaseModel{})
	result := new([]model.CaseModel)
	err := query.
		Preload("Contributor.Lawyer").
		Preload("Uploader.Lawyer").
		Preload("Client").
		Find(result).Error
	if err != nil {
		return nil, r.maskError(err)
	}
	return *result, nil
}
