package repository

import (
	"gh5-backend/internal/model/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleRepository struct {
	Repository[entity.RoleModel]
	Log *logrus.Logger
}

func NewRoleRepository(conn *gorm.DB, log *logrus.Logger) *RoleRepository {
	model := entity.RoleModel{}
	repository := NewRepository(conn, model, model.TableName())
	return &RoleRepository{
		Repository: *repository,
		Log:        log,
	}
}
