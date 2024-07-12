package repository

import (
	"context"
	"gh5-backend/internal/model/entity"
	model "gh5-backend/internal/model/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.UserModel]
	Log *logrus.Logger
}

func NewUserRepository(conn *gorm.DB, log *logrus.Logger) *UserRepository {
	model := entity.UserModel{}
	repository := NewRepository(conn, model, model.TableName())
	return &UserRepository{
		Repository: *repository,
		Log:        log,
	}
}

func (r *Repository[T]) FindByEmail(ctx context.Context, email string) (*model.UserModel, error) {
	query := r.getConn().Model(model.UserModel{})
	result := new(model.UserModel)
	err := query.Where("email", email).Preload("Role").First(result).Error
	if err != nil {
		return nil, r.maskError(err)
	}
	return result, nil
}

func (r *Repository[T]) UpdatesByID(ctx context.Context, id string, data *entity.CaseModel) (model.CaseModel, error) {
	query := r.getConn().Table(r.entityName)
	result := model.CaseModel{}
	err := query.Where("id", id).Updates(
		map[string]interface{}{
			"id":               data.ID,
			"case_number":      data.CaseNumber,
			"case_description": data.CaseDescription,
			"case_detail":      data.CaseDetail,
			"is_active":        data.IsActive,
			"status":           data.Status,
			"contributor_id":   data.ContributorID,
			"uploader_id":      data.UploaderID,
		},
	).Error
	if err != nil {
		return result, r.maskError(err)
	}
	return result, nil
}

func (r *Repository[T]) UpdateLawyerID(ctx context.Context, id string, lawyerID string) error {
	query := r.getConn().Table(r.entityName)
	err := query.Where("id", id).Update("lawyer_id", id).Error
	if err != nil {
		return r.maskError(err)
	}
	return nil
}
