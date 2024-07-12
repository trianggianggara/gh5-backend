package repository

import (
	"context"
	model "gh5-backend/internal/model/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VoteRepository struct {
	Repository[model.VoteModel]
	Log *logrus.Logger
}

func NewVoteRepository(conn *gorm.DB, log *logrus.Logger) *VoteRepository {
	model := model.VoteModel{}
	repository := NewRepository(conn, model, model.TableName())
	return &VoteRepository{
		Repository: *repository,
		Log:        log,
	}
}

func (r *VoteRepository) Revote(ctx context.Context, caseID string, userID string) (model.VoteModel, error) {
	r.checkTrx(ctx)
	query := r.getConn().Table(r.entityName)
	result := model.VoteModel{}
	err := query.Where("case_id = ? AND user_id = ?", caseID, userID).Update("is_active", true).Error
	if err != nil {
		return result, r.maskError(err)
	}
	return result, nil
}

func (r *VoteRepository) Downvote(ctx context.Context, caseID string, userID string) (model.VoteModel, error) {
	r.checkTrx(ctx)
	query := r.getConn().Table(r.entityName)
	result := model.VoteModel{}
	err := query.Where("case_id = ? AND user_id = ?", caseID, userID).Update("is_active", false).Error
	if err != nil {
		return result, r.maskError(err)
	}
	return result, nil
}

func (r *VoteRepository) VoteCountByCaseID(ctx context.Context, caseID string) (*model.VoteCount, error) {
	r.checkTrx(ctx)
	query := r.conn.Model(model.VoteCount{})
	result := new(model.VoteCount)
	err := query.Where("case_id = ?", caseID).Preload("Case").First(result).Error
	if err != nil {
		return nil, r.maskError(err)
	}
	return result, nil
}

func (r *VoteRepository) VoteCount(ctx context.Context) ([]model.VoteCount, error) {
	r.checkTrx(ctx)
	query := r.conn.Model(model.VoteCount{})
	result := new([]model.VoteCount)
	err := query.
		Preload("Case.Client.Role").
		Preload("Case.Uploader.Lawyer").
		Preload("Case.Uploader.Role").
		Preload("Case.Contributor.Lawyer").
		Preload("Case.Contributor.Role").
		Find(result).Error
	if err != nil {
		return nil, r.maskError(err)
	}
	return *result, nil
}

func (r *VoteRepository) FindVoteByCaseAndUserID(ctx context.Context, caseID string, userID string) (model.VoteModel, error) {
	r.checkTrx(ctx)
	query := r.getConn().Table(r.entityName)
	result := model.VoteModel{}
	err := query.Where("case_id = ? AND user_id = ?", caseID, userID).First(result).Error
	if err != nil {
		return result, r.maskError(err)
	}
	return result, nil
}
