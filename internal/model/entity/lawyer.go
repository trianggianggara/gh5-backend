package entity

import (
	"context"

	abstraction "gh5-backend/internal/model/base"
	constant "gh5-backend/pkg/constants"
	"gh5-backend/pkg/ctxval"
	"gh5-backend/pkg/utils/date"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LawyerEntity struct {
	Specialization string `json:"specialization" validate:"required"`
	UserID         string `json:"user_id" validate:"required" gorm:"not null"`
	IsActive       bool   `json:"is_active"`

	User *UserModel `json:"user" gorm:"foreignkey:UserID"`
}

type LawyerModel struct {
	abstraction.Entity
	LawyerEntity
	Context context.Context `json:"-" gorm:"-"`
}

func (LawyerModel) TableName() string {
	return "lawyers"
}

func (m *LawyerModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	m.IsActive = true
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *LawyerModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()

	authCtx := ctxval.GetAuthValue(m.Context)
	if authCtx != nil {
		m.ModifiedBy = &authCtx.Name
	}
	return
}
