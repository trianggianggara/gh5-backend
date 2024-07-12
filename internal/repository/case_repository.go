package repository

import (
	"context"
	"gh5-backend/internal/model/entity"
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

func (r *CaseRepository) FindAll(ctx context.Context) ([]model.CaseModel, error) {
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

func (r *CaseRepository) FindCaseByUserID(ctx context.Context, userID string) ([]model.CaseDetails, error) {
	r.checkTrx(ctx)
	query := r.getConn().Model(model.CaseDetails{})
	result := new([]model.CaseDetails)
	err := query.
		Where("user_id = ?", userID).
		Preload("Contributor.Lawyer").
		Preload("Uploader.Lawyer").
		Preload("Client").
		Find(result).Error
	if err != nil {
		return nil, r.maskError(err)
	}
	return *result, nil
}

func (r *CaseRepository) FindCaseByLawyerID(ctx context.Context, lawyerID string, status string) ([]model.LawyerCase, error) {
	r.checkTrx(ctx)
	query := r.getConn().Model(model.LawyerCase{})
	result := new([]model.LawyerCase)
	err := query.
		Where("lawyer_id = ? and status = ?", lawyerID, status).
		Preload("Contributor.Lawyer").
		Preload("Uploader.Lawyer").
		Preload("Client").
		Find(result).Error
	if err != nil {
		return nil, r.maskError(err)
	}
	return *result, nil
}

func (r *CaseRepository) UpdatesByID(ctx context.Context, id string, data *entity.CaseModel) (model.CaseModel, error) {
	r.checkTrx(ctx)
	query := r.getConn().Table(r.entityName)
	result := model.CaseModel{}
	err := query.Where("id", id).Updates(
		map[string]interface{}{
			"id":               data.ID,
			"case_number":      data.CaseNumber,
			"case_description": data.CaseDescription,
			"case_detail":      data.CaseDetail,
			"document":         data.Document,
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
