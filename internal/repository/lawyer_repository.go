package repository

import (
	"context"
	"gh5-backend/internal/model/entity"
	model "gh5-backend/internal/model/entity"

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

func (r *LawyerRepository) UpdatesByID(ctx context.Context, id string, data *model.LawyerModel) (model.LawyerModel, error) {
	query := r.getConn().Table(r.entityName)
	result := model.LawyerModel{}
	err := query.Where("id", id).Updates(
		map[string]interface{}{
			"position":       data.Position,
			"specialization": data.Specialization,
		},
	).Error
	if err != nil {
		return result, r.maskError(err)
	}
	return result, nil
}
