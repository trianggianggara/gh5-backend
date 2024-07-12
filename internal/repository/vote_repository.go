package repository

import (
	"context"
	"gh5-backend/internal/model/entity"
	model "gh5-backend/internal/model/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VoteRepository struct {
	Repository[entity.VoteModel]
	Log *logrus.Logger
}

func NewVoteRepository(conn *gorm.DB, log *logrus.Logger) *VoteRepository {
	model := entity.VoteModel{}
	repository := NewRepository(conn, model, model.TableName())
	return &VoteRepository{
		Repository: *repository,
		Log:        log,
	}
}

func (r *Repository[T]) Downvote(ctx context.Context, caseID string, userID string) (model.VoteModel, error) {
	query := r.getConn().Model(model.VoteModel{})
	result := model.VoteModel{}
	err := query.Where("case_id = ? AND user_id = ?", caseID, userID).Update("is_active", false).Error
	if err != nil {
		return result, r.maskError(err)
	}
	return result, nil
}
