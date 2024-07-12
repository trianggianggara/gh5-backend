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

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.UserModel, error) {
	r.checkTrx(ctx)
	query := r.getConn().Model(model.UserModel{})
	result := new(model.UserModel)
	err := query.Where("email", email).Preload("Role").First(result).Error
	if err != nil {
		return nil, r.maskError(err)
	}
	return result, nil
}

func (r *UserRepository) UpdateLawyerID(ctx context.Context, id string, lawyerID string) error {
	r.checkTrx(ctx)
	query := r.getConn().Table(r.entityName)
	err := query.Where("id", id).Update("lawyer_id", lawyerID).Error
	if err != nil {
		return r.maskError(err)
	}
	return nil
}
